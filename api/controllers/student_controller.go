package controllers

import (
	"database/sql"
	"fmt"
	"proj_struct/api/models"
	"proj_struct/internal/pkg1"
	"net/http"
	"strconv"
)

func GetStudentByID(id int) (*models.Student, error) {
	// Connect to the database
	db, err := pkg1.SetupDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the student by ID
	query := "SELECT id, name, grade FROM students WHERE id = ?"
	row := db.QueryRow(query, id)

	var student models.Student
	err = row.Scan(&student.ID, &student.Name, &student.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Student with ID %d not found", id)
		}
		return nil, err
	}

	return &student, nil
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	// Extract student ID from request parameters
	idParam := r.URL.Query().Get("id")
	studentID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	// Get student details
	student, err := GetStudentByID(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Return student details as JSON
	// (In a real application, you would use a JSON library like encoding/json)
	fmt.Fprintf(w, "Student Details: %+v", student)
}
