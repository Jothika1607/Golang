package main

import (
	"fmt"
	"proj_struct/api/controllers"
	"proj_struct/internal/pkg1"
	"net/http"
)

func main() {
	fmt.Println("Main application")

	// Example usage of controllers and utilities
	student, err := controllers.GetStudentByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Student Details: %+v\n", student)

	pkg1.StringUtilFunction()

	// Start a simple HTTP server
	http.HandleFunc("/student", controllers.StudentHandler)
	http.ListenAndServe(":8080", nil)
}
