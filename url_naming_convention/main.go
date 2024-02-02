//This program show the rest url api naming convention in go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Structs are named using CamelCase.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "name1"},
	{ID: 2, Name: "name2"},
}

// Function names are in CamelCase.
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Variables are named either in camelCase, PascalCase or SnakeCase.
	vars := mux.Vars(r)
	userID := vars["id"]

	id := parseID(userID)
	if id == -1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid user ID")
		return
	}

	var foundUser *User
	for _, user := range users {
		if user.ID == id {
			foundUser = &user
			break
		}
	}

	if foundUser != nil {
		json.NewEncoder(w).Encode(foundUser)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "User not found")
	}
}

func parseID(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1
	}
	return id
}

func main() {
	router := mux.NewRouter()

	// Define API routes
	//Use plural nouns to represent resources in the URL.
	//Use HTTP methods (GET, POST, PUT, DELETE) to represent actions on resources.
	//Use hyphens to separate words in URLs for better readability.
	//Include API versioning in the URL to ensure backward compatibility.
	router.HandleFunc("/v1/users", getAllUsersHandler).Methods("GET")
	//Nested Resources for Relationships
	//Use nested resources to represent hierarchical relationships.
	router.HandleFunc("/v1/users/{id}", getUserByIDHandler).Methods("GET")

	port := 8080
	fmt.Printf("Starting server on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
