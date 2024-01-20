package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go delay(c)
	<-c
}
func delay(c chan int) {
	time.Sleep(3 * time.Millisecond)
	fmt.Println("Channel creation")
	c <- 100
}
