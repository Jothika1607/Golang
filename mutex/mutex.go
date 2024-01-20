package main

import (
	"fmt"
	"sync"
)

const n = 10

func main() {
	num := 10000
	mutex := sync.Mutex{}
	ch := make(chan struct{})

	go func() {
		for i := 0; i < n; i++ {
			mutex.Lock()
			num++
			mutex.Unlock()
		}
		ch <- struct{}{}
	}()

	go func() {
		for i := 0; i < n; i++ { // Corrected the loop
			mutex.Lock()
			num--
			mutex.Unlock()
		}
		ch <- struct{}{}
	}()

	<-ch
	<-ch
	fmt.Println(num)
}
