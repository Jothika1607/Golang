package main

import (
	"fmt"
)

var n2 int = 20 //globle variable
func main() {
	var n1 int = 10 // local variable and access inside the function
	sum := n1 + n2
	fmt.Printf("Sum=%d", sum)
}
