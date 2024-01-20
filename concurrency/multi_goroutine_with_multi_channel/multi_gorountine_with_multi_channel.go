package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels of type string
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines, each sending a message to its respective channel
	go sendMessage("Hello from Goroutine 1!", ch1)
	go sendMessage("Greetings from Goroutine 2!", ch2)

	// Main goroutine receives and prints messages from both channels
	receivedMessage1 := <-ch1
	receivedMessage2 := <-ch2

	fmt.Println("Main goroutine received:", receivedMessage1)
	fmt.Println("Main goroutine received:", receivedMessage2)
}

// Function to send a message to the channel
func sendMessage(message string, ch chan string) {
	// Simulate some processing time
	time.Sleep(2 * time.Second)

	// Send a message to the channel
	ch <- message
}
