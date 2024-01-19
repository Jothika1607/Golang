package main

import (
	"fmt"
	"net/http"
)

// Handler function for the "/hello" path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is a handler function!")
}

func main() {
	// Registering the helloHandler for the "/hello" path
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
