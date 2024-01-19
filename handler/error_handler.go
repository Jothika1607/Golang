package main

import (
	"fmt"
	"net/http"
)

// Handler function for the "/hello" path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is a handler function!")
}	

// Error handler function
func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
}

func main() {
	// Registering the helloHandler for the "/hello" path
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Simulating an error for demonstration purposes
		err := fmt.Errorf("something went wrong")
		if err != nil {
			// Call the error handler with the response writer, request, and error
			errorHandler(w, r, err)
			return
		}

		// If no error occurred, proceed with the regular handler logic
		helloHandler(w, r)
	})

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
