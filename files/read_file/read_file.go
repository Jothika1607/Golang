package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	_, err = file.WriteString("Welcome to Tutorials Point")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}
func ReadFile() {
	fileName := "test.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\n%s", data)
}
func main() {
	CreateFile()
	ReadFile()
}
