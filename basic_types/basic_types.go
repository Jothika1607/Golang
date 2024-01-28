package main

import "fmt"

func main() {
	//Numeric datatypes
	var num1 int = 42
	var num2 float64 = 3.14
	fmt.Println("Integer:", num1)
	fmt.Println("Float:", num2)

	//string datatype
	message := "Hello, Golang!"
	fmt.Println(message)

	//Boolean datatype
	isTrue := true
	isFalse := false
	fmt.Println("True:", isTrue)
	fmt.Println("False:", isFalse)

	//pointer datatype
	var x int = 10
	var ptr *int = &x
	fmt.Println("Value of x:", x)
	fmt.Println("Value through pointer:", *ptr)
}
