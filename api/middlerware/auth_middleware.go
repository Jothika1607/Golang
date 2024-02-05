package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement authentication logic here
		// For simplicity, allowing all requests
		fmt.Println("Authentication passed")
		next.ServeHTTP(w, r)
	})
}
