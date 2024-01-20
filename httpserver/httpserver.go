package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function for incoming HTTP requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Write the response body
		fmt.Fprintf(w, "Hello, This is a  http server concept")
	}

	// Register the handler function for a specific route ("/" in this case)
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
