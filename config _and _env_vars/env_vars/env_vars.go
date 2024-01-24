// main.go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    int
	Debug   bool
}

func loadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	portStr := os.Getenv("APP_PORT")
	if portStr == "" {
		return nil, fmt.Errorf("APP_PORT environment variable is not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse APP_PORT: %v", err)
	}

	debugStr := os.Getenv("DEBUG")
	if debugStr == "" {
		return nil, fmt.Errorf("DEBUG environment variable is not set")
	}

	debug, err := strconv.ParseBool(debugStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DEBUG: %v", err)
	}

	config := &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    port,
		Debug:   debug,
	}

	return config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug Mode: %v\n", config.Debug)
}
