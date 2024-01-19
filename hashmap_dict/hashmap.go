package main

import "fmt"

func main() {

	// Initialize an empty map with string keys and integer values
	hashmap := make(map[string]int)

	// Add some items to the map
	hashmap["pencil"] = 10
	hashmap["pen"] = 20
	hashmap["scale"] = 15

	// Print the entire map
	fmt.Println("The hashmap created above is: ")
	fmt.Println(hashmap)

	// Access a specific item by its key
	fmt.Println("The value of specific element from hashmap is:")
	fmt.Println(hashmap["pencil"])

	// Update the value of an existing item
	hashmap["scale"] = 20

	// Delete an item from the map
	delete(hashmap, "pencil")

	// Print the updated map
	fmt.Println("The hashmap after deleting an item from it is:")
	fmt.Println(hashmap)
}
