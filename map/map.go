package main

import "fmt"

func main() {

	//create map
	a := map[string]int{"n1": 20, "n2": 30}
	fmt.Println(a)

	fmt.Println(a["n1"]) // to access the value of the key

	a["n2"] = 50 //Change the value
	fmt.Println(a)

	a["n3"] = 100 // add keys and values
	a["n4"] = 200
	fmt.Println(a)

	delete(a, "n3") //delete from map with keys
	fmt.Println(a)

	for index, item := range a { //iterate using loop for map
		fmt.Printf("index of %s: %d \n", index, item)
	}
}
