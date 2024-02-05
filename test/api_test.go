package test

import (
	"myproject/api/controllers"
	"testing"
)

func TestGetStudentByID(t *testing.T) {
	// Implement test cases for GetStudentByID function
	student := controllers.GetStudentByID(1)
	if student == nil {
		t.Error("Expected non-nil student, got nil")
	}
}
