package main

import (
	"context"
	"fmt"
)

func myfunc(ctx context.Context) {
	fmt.Println("Go language")
}

func main() {
	ctx := context.TODO()
	myfunc(ctx)
}

