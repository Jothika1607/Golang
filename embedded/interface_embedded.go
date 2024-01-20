package main

import "fmt"

// Shape interface
type Shape interface {
	area() float64
}

// Rectangle struct
type Rectangle struct {
	length, width float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Embedding Shape interface in Rectangle struct
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// Embedding Shape interface in Circle struct
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// Creating objects of Rectangle and Circle
	r := Rectangle{length: 5, width: 3}
	c := Circle{radius: 4}

	// Creating slice of Shape interface type and adding objects to it
	shapes := []Shape{r, c}

	// Iterating over slice and calling area method on each object
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
