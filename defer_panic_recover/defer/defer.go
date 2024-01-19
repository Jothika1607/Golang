package main

import "fmt"

func main() {
	// defer the execution of Println() function
	defer fmt.Println("Three")
	fmt.Println("One")
	fmt.Println("Two")

	//multiple defer
	defer fmt.Println("defer1")
	defer fmt.Println("Defer2")
	defer fmt.Println("Defer3")
}
