package main

import (
	"fmt"
	"time"
)

func rountine(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Printf("%s\n", s)
	}
}
func main() {
	go rountine("Hello")
	rountine("World")
}
