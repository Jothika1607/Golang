package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := "SGVsbG8sIHdvcmxkIQ=="
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding string:", err)
		return
	}
	fmt.Println(string(decoded))
}
