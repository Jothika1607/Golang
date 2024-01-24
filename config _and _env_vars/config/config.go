// main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
	Debug   bool   `json:"debug"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug Mode: %v\n", config.Debug)
}
