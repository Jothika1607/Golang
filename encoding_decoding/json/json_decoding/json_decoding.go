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
	// JSON string representing a Person
	jsonString := `{"name":"Alice","age":25,"email":"alice@example.com"}`

	// Creating a Person struct to hold the decoded data
	var person Person

	// Decoding (unmarshalling) the JSON string into the struct
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Printing the decoded data
	fmt.Printf("Name: %s, Age: %d, Email: %s\n", person.Name, person.Age, person.Email)
}
