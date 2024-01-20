package main

import (
	"fmt"
	"sync"
	"time"
)

var resource int32 = 0

func Read(rwm *sync.RWMutex, wg *sync.WaitGroup) {

	rwm.RLock()
	fmt.Println("read lock acquired")

	time.Sleep(time.Second * 3)

	fmt.Println("read lock released")
	rwm.RUnlock()
	wg.Done()
}

func Write(rwm *sync.RWMutex, wg *sync.WaitGroup) {

	rwm.Lock()
	fmt.Println("write lock acquired")

	resource = resource + 1
	time.Sleep(time.Second * 3)

	fmt.Println("write lock released")
	rwm.Unlock()
	wg.Done()

}

func main() {
	rwm := &sync.RWMutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go Write(rwm, wg)
		go Read(rwm, wg)
	}

	wg.Wait()
}
