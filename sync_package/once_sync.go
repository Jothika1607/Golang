package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	RunOnce := func() {
		fmt.Println("Run once")
	}
	done := make(chan string)
	for i := 0; i < 5; i++ {
		go func() {
			once.Do(RunOnce)
			done <- "Hi"
		}()
	}
	for i := 0; i < 5; i++ {
		<-done
	}
}
