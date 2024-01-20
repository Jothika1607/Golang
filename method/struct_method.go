package main

import (
	"fmt"
)

type student struct { //student structure
	fname string
	lname string
	id    int
	age   int
}

func (s student) display() {
	fmt.Println("First name :", s.fname)
	fmt.Println("Last name :", s.lname)
	fmt.Println("Id :", s.id)
	fmt.Println("Age :", s.age)
}
func main() {
	std := student{
		fname: "Jothika",
		lname: "Sooludiayar",
		id:    2100,
		age:   23,
	}
	std.display()
}
