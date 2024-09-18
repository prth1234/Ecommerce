package storeapi

import (
	"context"
	"net/http"
	"store/jwthelper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		claims, err := jwthelper.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add the username to the request context
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
