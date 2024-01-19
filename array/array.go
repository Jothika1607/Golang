package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}    //type is defined using var keyword
	var a1 = [...]int{1, 2, 3, 4} ////type is inferred using var keyword
	arr2 := [5]int{4, 5, 6, 7, 8} //type is defined using :=
	a2 := [...]int{4, 5, 6}       //type is inferred using :=

	fmt.Println(arr1, a1) // access the array
	fmt.Println(arr2, a2)
	fmt.Println(arr1[2]) /// access the array using index

	a1[1] = 20 //change the value in array
	fmt.Println(a1)

	//Array initialization
	i := [5]int{}              //not initialized
	j := [5]int{9, 8}          //partially initialized
	k := [5]int{7, 6, 5, 4, 3} //not initialized
	fmt.Println(i, j, k)

	var n1 = [5]int{1: 20, 3: 50}
	fmt.Println(n1)

	n := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(len(n))
}
