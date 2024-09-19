package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	storeapi "store"
	"store/gen/store"
	"sync"
	"syscall"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Define command line flags

	// Serve Redoc static files
	fs := http.FileServer(http.Dir("./redoc"))

	// Serve the Redoc UI
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	// Serve the OpenAPI spec (openapi.json)
	http.Handle("/openapi.json", http.FileServer(http.Dir("./path/to/openapi")))

	// Serve the Redoc HTML page that loads the OpenAPI spec
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
			  <head>
				<title>Redoc</title>
				<!-- Needed for adaptive design -->
				<meta charset="utf-8"/>
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<!--
				Redoc script
				-->
				<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
			  </head>
			  <body>
				<redoc spec-url='/openapi.json'></redoc>
			  </body>
			</html>
		`))
	})

	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "8000", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		dbURLF    = flag.String("db-url", "postgres://postgres:kundan@localhost:5432/EcommerceDemo?sslmode=disable", "PostgreSQL connection URL")
	)
	flag.Parse()

	// Setup logger
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[store] ", log.Ltime)
	}

	// Initialize the services
	var (
		storeSvc store.Service
	)
	{
		db, err := sql.Open("postgres", *dbURLF)
		if err != nil {
			logger.Fatalf("Failed to connect to database: %v", err)
		}
		defer db.Close()

		storeSvc = storeapi.NewStore(db)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		storeEndpoints *store.Endpoints
	)
	{
		storeEndpoints = store.NewEndpoints(storeSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:" + *httpPortF
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s", addr, err)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, storeEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: localhost)", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Printf("exited")
}
