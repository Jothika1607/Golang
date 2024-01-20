package main

import "fmt"

type Sum int // Type definition
func (s1 Sum) add(s2 Sum) Sum { // Defining a method with non-struct type receiver
	return s1 + s2
}

/*if you try to run this code, then compiler will throw an error
func(d1 int)add(d2 int)int{
return d1 + d2
}
*/
func main() { // Main function
	value1 := Sum(100)
	value2 := Sum(200)
	res := value1.add(value2)
	fmt.Println("Final result: ", res)
}
