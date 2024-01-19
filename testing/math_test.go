package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Use the testing type *testing.T to run tests and report results.
	result := Add(2, 3)
	expected := 5

	// Check if the result is equal to the expected value.
	if result != expected {
		// If not equal, report a test failure using t.Errorf.
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
