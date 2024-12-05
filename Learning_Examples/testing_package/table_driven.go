package main

import "testing"

// Function to test
func Multiply(a, b int) int {
	return a * b
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 4, -8},
	}

	for _, tt := range tests {
		result := Multiply(tt.a, tt.b)
		if result != tt.expect {
			t.Errorf("Multiply(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expect, result)
		}
	}
}
