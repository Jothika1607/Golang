package main

import (
	"fmt"
	"sync"
)

const n = 10

type Counter struct {
	count int
	sync.Mutex
}

func main() {
	counter := Counter{count: 100}
	ch := make(chan struct{})

	go func() {
		for i := 0; i < n; i++ {
			counter.Lock()
			counter.count++
			counter.Unlock()
		}
		ch <- struct{}{}
	}()

	go func() {
		for i := 0; i < n; i++ { // Corrected the loop
			counter.Lock()
			counter.count--
			counter.Unlock()
		}
		ch <- struct{}{}
	}()

	<-ch
	<-ch
	fmt.Println(counter.count)
}
