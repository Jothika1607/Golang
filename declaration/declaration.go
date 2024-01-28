package main

import (
	"fmt"
)

var a string = "JO" // type is string
var b = 20          // type is inferred
var d bool          // declaration without value
var str string
var x, y, z int = 20, 40, 50 //multiple variable declaration
var i, j = 100, "str"
var ( //varibale declaration within block
	n1 int
	n2 int    = 1
	n3 string = "hello"
)

const P = 10	//constant variable declaration

func main() {
	c := 3.5 // tpye is inferred
	m, n := 200, "str1"
	str = "Jothika"
	fmt.Println(P)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Print(x, " ", y, z)
	fmt.Println(m, n)
	fmt.Println(i, j)
	fmt.Println(str)
	fmt.Println(n1, n2, n3)
	fmt.Println("Hello World!")
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)

	var name string

	// takes input value for name
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Name: %s", name) 
}
