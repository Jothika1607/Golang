package main

import (
	"fmt"
	"example.com/myproject/api"
	"example.com/myproject/internal/pkg1"
	"example.com/myproject/internal/pkg2"
)

func main() {
	fmt.Println("Hello from myapp!")

	api.HandleAPI()    // Call API-related logic
	pkg1.DoSomething() // Call internal package logic
	pkg2.DoSomething() // Call internal package logic
}
