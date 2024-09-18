package main

import (
	"context"
	"fmt"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"log"
	"net/http"
	"net/url"
	"os"
	storesvr "store/gen/http/store/server"
	store "store/gen/store"
	custommiddleware "store/middleware"
	"sync"
	"time"
)

// CustomLogger implements the middleware.Logger interface
type CustomLogger struct {
	logger *log.Logger
}

func (l *CustomLogger) Log(keyvals ...interface{}) error {
	l.logger.Println(keyvals...)
	return nil
}

// Define a custom key type for the request ID
type requestIDKey struct{}

func handleHTTPServer(ctx context.Context, u *url.URL, endpoints *store.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {
	// Setup custom logger adapter
	customLogger := &CustomLogger{logger: logger}

	// Provide the transport specific request decoder and response encoder.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers.
	var (
		storeServer *storesvr.Server
	)
	{
		eh := errorHandler(logger)
		storeServer = storesvr.New(endpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				storeServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	storesvr.Mount(mux, storeServer)

	// Wrap the multiplexer with additional middlewares.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(customLogger)(handler)
		handler = httpmdlwr.RequestID()(handler)

		// Add custom middleware to store request ID in context
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			id := r.Header.Get("X-Request-Id")
			if id == "" {
				id = "unknown"
			}
			ctx = context.WithValue(ctx, requestIDKey{}, id)
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		})

		// Add JWT middleware to all routes except login and create user
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/login" && r.URL.Path != "/users" {
				custommiddleware.JWTAuthMiddleware(mux).ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		})
	}

	// Start HTTP server using default configuration.
	server := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range storeServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- server.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		// Retrieve the request ID from the context
		id, _ := ctx.Value(requestIDKey{}).(string)
		if id == "" {
			id = "unknown"
		}
		_, _ = w.Write([]byte(fmt.Sprintf("[%s] encoding: %s", id, err.Error())))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
