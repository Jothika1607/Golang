package main

import (
	"fmt"
)

func main() {
	s1 := []string{"Go", "Slices", "Are", "Powerful"} //declare the slice
	fmt.Println(s1)
	fmt.Println(len(s1)) //length of the slice
	fmt.Println(cap(s1)) //capacity of the slice

	//create slice from array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	//create slice using make()
	myslice1 := make([]int, 5, 10) //with capacity
	fmt.Println(myslice1)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

	myslice2 := make([]int, 5) //with  omitted capacity
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))

	slice := []int{2, 3, 4}
	fmt.Println(slice[1]) // access the slice using index

	slice[1] = 20
	fmt.Println(slice[1]) // chnage the slice value using index

	myslice2 = append(myslice2, 20, 21) // append
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))

	sl1 := []int{1, 2, 3}
	sl2 := []int{4, 5, 6}
	sl3 := append(sl1, sl2...) // append one slice to another

	fmt.Printf("myslice3=%v\n", sl3)
	fmt.Printf("length=%d\n", len(sl3))
	fmt.Printf("capacity=%d\n", cap(sl3))

	//memory efficiency
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
