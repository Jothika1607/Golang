package main

import (
	"log"
	"time"
)

func main() {
	// Log a message at the beginning
	log.Println("Program started")

	// Simulate some processing time
	doSomeWork()

	// Introduce a sleep of 2 seconds
	time.Sleep(2 * time.Second)
	log.Println("Slept for 2 seconds")

	// Log a message at the end
	log.Println("Program completed")
}

func doSomeWork() {
	// Simulate some processing time
	log.Println("Doing some work...")
	time.Sleep(1 * time.Second)
	log.Println("Work completed")
}
