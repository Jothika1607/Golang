// Go program to illustrate the
// concept of variadic function
package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {

	// zero argument
	fmt.Println(joinstr())

	// multiple arguments
	fmt.Println(joinstr("Jothika", "Sooludaiyar"))
	fmt.Println(joinstr("Jo", "Soolu", "daiyar"))
	fmt.Println(joinstr("J", "o", "t", "h", "i"))

}
