package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Printf("Worker number: %d running\n", x)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
