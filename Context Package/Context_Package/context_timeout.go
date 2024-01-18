package main

import (
	"context"
	"fmt"
	"time"
)

func myfunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("Processing")
		}
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go myfunc(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}
