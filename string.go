package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go Language"          //declare :=
	str1 := `Jothika sooludaiyar` //declare backtick ``
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Printf("%c", str[0]) //access individual character of string

	// create string
	message := "Welcome to Programiz"
	fmt.Println(message)
	fmt.Println(len(message))
	// use len() function to count length
	stringLength := len(message)
	fmt.Println(stringLength)

	m1 := "I am"
	m2 := " Jothika Sooludiayar"
	display := m1 + " " + m2 //Join two strings
	fmt.Println(display)

	text := "Go program and language"
	c1 := "program"
	c2 := "programmiz"
	c3 := "program"
	fmt.Println(strings.Compare(c1, c2)) //compare() function
	fmt.Println(strings.Compare(c1, c3))

	fmt.Println(strings.Contains(text, c1)) //contains() check if substring in string

	t1 := "car"
	replace := strings.Replace(t1, "r", "t", 1) //replace of character
	fmt.Println(replace)

	x := "jothika sooludaiyar"
	x1 := strings.ToUpper(x) //uppercase
	fmt.Println(x1)

	y := "JOTHIKA SOOLUDAIYAR"
	y1 := strings.ToLower(y) //lowercase
	fmt.Println(y1)

	x2 := "enjoy th life"
	smsg := strings.Split(x2, " ") //split()
	fmt.Println(smsg)

	msg := "This is a \" GO \" language" // escape sequence
	fmt.Println(msg)
}
