package main

import "fmt"

func main() {
	a := 10
	b := 3

	//Arithmetic operators
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder:", a%b)

	//comparison operators
	fmt.Println("Equal to:", a == b)
	fmt.Println("Not equal to:", a != b)
	fmt.Println("Less than:", a < b)
	fmt.Println("Greater than:", a > b)
	fmt.Println("Less than or equal to:", a <= b)
	fmt.Println("Greater than or equal to:", a >= b)

	//Logical operators
	a1 := true
	b1 := false
	fmt.Println("Logical AND:", a1 && b1)
	fmt.Println("Logical OR:", a1 || b1)
	fmt.Println("Logical NOT:", !a1)

	//Bitwise operators
	x := 5
	y := 3
	fmt.Printf("Bitwise AND: %08b\n", x&y)
	fmt.Printf("Bitwise OR: %08b\n", x|y)
	fmt.Printf("Bitwise XOR: %08b\n", x^y)
	fmt.Printf("Left shift: %08b\n", x<<1)
	fmt.Printf("Right shift: %08b\n", x>>1)

	//Assignment operator
	m := 10

	// Simple assignment
	n := m

	// Compound assignment
	m += 5
	fmt.Println("Simple Assignment:", n)
	fmt.Println("Compound Assignment:", m)
}
