// main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
	}

	for _, tc := range cases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
