package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//Using fmt package
	var num int
	fmt.Println("Get input:")
	fmt.Scanf("%d", &num)
	fmt.Printf("Print number:%d", num)

	//Using math package
	fmt.Println(math.Sqrt(5))        //Square root
	fmt.Println(math.Cbrt(5))        //cube root
	fmt.Println(math.Min(678, 3758)) //Minimum
	fmt.Println(math.Max(678, 3758)) //Maximum
	fmt.Println(math.Mod(100, 4))    //remainder

	//Using strings package
	// create a string array
	stringArray := []string{"I love", "Go Programming"}

	// join elements of array with space in between
	joinedString := strings.Join(stringArray, " ")
	fmt.Println(joinedString)
}
