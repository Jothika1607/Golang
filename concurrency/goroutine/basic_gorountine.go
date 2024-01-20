package main

import (
	"fmt"
	"time"
)

func delay(s string) {
	time.Sleep(3 * time.Millisecond)
	fmt.Printf("%s\n", s)
}
func main() {
	go delay("First")
	go delay("Second")
	go delay("Third")
	time.Sleep(3 * time.Second)
}
