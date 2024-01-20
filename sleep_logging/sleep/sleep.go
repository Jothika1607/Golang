package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of the program")

	// Sleep for 2 seconds
	time.Sleep(2 * time.Second)

	fmt.Println("After sleeping for 2 seconds")

	// Sleep for 500 milliseconds
	time.Sleep(500 * time.Millisecond)

	fmt.Println("After sleeping for 500 milliseconds")
}
