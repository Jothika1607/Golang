package main

import "fmt"

func main() {

	age := -14

	// create an error using Efforf()
	error := fmt.Errorf("%d is negative\nAge can't be negative", age)

	if age < 0 {
		fmt.Print(error)
	} else {
		fmt.Printf("Age: %d", age)
	}
}
