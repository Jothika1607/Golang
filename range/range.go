package main

import "fmt"

func main() {
	n := [5]int{10, 20, 30, 40}
	fmt.Println(n)

	for index, item := range n { //range with array
		fmt.Printf("Index %d = %d\n", index, item)
	}

	str := "Golang"
	fmt.Println(str)
	for index, item := range str { //range with string
		fmt.Printf("Index %d=%c\n", index, item)
	}

	m := map[string]string{"name": "JO", "Age": "20", "Id": "234522"}
	fmt.Println(m)
	for index, item := range m { //range with map
		fmt.Printf("Index %s=%s\n", index, item)
	}

	for index := range m { //access key of map using range
		fmt.Println(index)
	}
}
