package main

import (
	"fmt"
)

type student struct {
	Name string
	id   int
}

type details struct {
	age     int
	student //embedded struct
}

func main() {
	d := details{
		age: 23,
		student: student{
			Name: "Jothika",
			id:   1230,
		},
	}

	fmt.Println("Name:", d.Name)
	fmt.Println("ID:", d.age)
	fmt.Println("Age:", d.id)
}
