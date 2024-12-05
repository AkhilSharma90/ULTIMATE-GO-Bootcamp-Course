package main

import "testing"

func TestOperations(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		result := Add(2, 3)
		if result != 5 {
			t.Errorf("Expected 5, got %d", result)
		}
	})

	t.Run("Multiplication", func(t *testing.T) {
		result := Multiply(2, 3)
		if result != 6 {
			t.Errorf("Expected 6, got %d", result)
		}
	})
}
