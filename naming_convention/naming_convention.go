package main

import (
	"fmt"
	"net/http"
)

// Constants are usually named in uppercase with underscores for readability.
const BASE_URL = "https://example.com/api"

// Structs are named using CamelCase.
type ItemDetails struct {
	ID    int
	Name  string
	Price float64
}

// Function names are in CamelCase.
func FetchDataFromAPI() {
	// Variables are named either in camelCase, PascalCase or SnakeCase.
	apiURL := BASE_URL + "/items"

	// Methods and functions should have descriptive names.
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer response.Body.Close()

	// Struct instances should be named with camelCase.
	var item ItemDetails

	// Use clear and concise variable names.
	// Loop through response, parse JSON, and populate the struct.
	// ...

	// Methods on structs are named with CamelCase as well.
	item.DisplayDetails()
}

// Methods on structs are named with CamelCase.
func (item *ItemDetails) DisplayDetails() {
	fmt.Printf("Item ID: %d\n", item.ID)
	fmt.Printf("Item Name: %s\n", item.Name)
	fmt.Printf("Item Price: $%.2f\n", item.Price)
}

func main() {
	FetchDataFromAPI()
}
