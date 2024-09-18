package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"store/jwthelper"
)

// JWTAuthMiddleware is the middleware that protects routes requiring authentication
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := bearerToken[1]
		claims, err := jwthelper.VerifyJWT(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		// Set the username in the context
		ctx := context.WithValue(r.Context(), "user", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
