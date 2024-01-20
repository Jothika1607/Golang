package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	// Creating a Person struct
	person := Person{Name: "John Doe", Age: 30}

	// Encoding (marshalling) the struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Printing the JSON data as a string
	fmt.Println(string(jsonData))
}
