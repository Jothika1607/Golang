package main

import (
	"fmt"
	"net/http"
)

// CustomHandler implements the http.Handler interface
type CustomHandler struct{}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a custom handler!")
}

func main() {
	// Create an instance of CustomHandler
	customHandler := &CustomHandler{}

	// Use customHandler as the handler for the root ("/") path
	http.Handle("/", customHandler)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
