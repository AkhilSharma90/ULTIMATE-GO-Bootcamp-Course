package main

import (
	"fmt"
	"testing"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func TestDivide(t *testing.T) {
	_, err := Divide(4, 0)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	result, err := Divide(10, 2)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
