package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Hello, world!"
	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encoded)
}
