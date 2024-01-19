package main

import (
	"fmt"
	"net/http"
)

// Middleware function
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Main handler
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is the main handler!")
}

func main() {
	// Wrap the main handler with the loggerMiddleware
	http.Handle("/hello", loggerMiddleware(http.HandlerFunc(mainHandler)))

	http.ListenAndServe(":8080", nil)
}
