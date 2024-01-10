package main

import (
	"testing"
)

// test function
func TestReturnGeeks(t *testing.T) {
	actualString := ReturnGeeks()
	expectedString := "go"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
