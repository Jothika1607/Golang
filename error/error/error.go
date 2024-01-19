package main

import (
	"errors" // import the errors package
	"fmt"
)

func main() {

	message := "Good day"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Nice day" {
		fmt.Println(myError)
	}

}
