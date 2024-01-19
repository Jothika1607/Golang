package main

import (
	"fmt"
)

func greet() { // define a function
	fmt.Println("Hello Jothika")
}

func sum(n1 int, n2 int) { // define a function with parameters
	sum := n1 + n2
	fmt.Printf("Sum : %d", sum)
}

func add(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func answer(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	diff := n1 - n2
	return sum, diff
}
func main() {
	greet()     // call a function
	sum(10, 20) //call a function with parameters
	result := add(50, 100)
	fmt.Printf("Sum=%d", result) // call a return value
	sum, diff := answer(100, 50)
	fmt.Printf("sum=%d,diff=%d", sum, diff)
}
