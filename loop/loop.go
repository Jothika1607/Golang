package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue //coninue
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break //break
		}
		fmt.Println(i)
	}
}
