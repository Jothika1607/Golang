package main

import (
	"log"
)

func main() {
	log.Print("This is a log message")
	log.Printf("Logging with format: %s", "Hello, Go!")
	log.Println("This is a log message with a newline")

	// Uncomment the line below to see how Fatal functions work
	log.Fatal("Fatal error: program cannot continue")
}
