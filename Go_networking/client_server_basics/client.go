package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	fmt.Fprintf(conn, "Hello, server!")

	// Receive and handle the server's response
	// (In a real application, you'd implement your specific logic here)
}
