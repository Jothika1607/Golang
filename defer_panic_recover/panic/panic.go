package main

import (
	"fmt"
)

func div(n1, n2 int) {
	if n2 == 0 {
		panic("It can't be work")
	} else {
		result := n1 / n2
		fmt.Println(result)
	}
}
func main() {
	div(10, 2)
	div(10, 5)
	div(10, 0)
	div(10, 3)
}
