package main

import (
	"errors"
	"fmt"
)

// Functions for arithmetic operations
func add(a, b float64) float64 {
	return a + b
}
func subtract(a, b float64) float64 {
	return a - b
}
func multiply(a, b float64) float64 {
	return a * b
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function to handle the calculation based on operator
func calculate(a, b float64, operator string) {
	var result float64
	var err error
	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result, err = divide(a, b)
	default:
		fmt.Println("Invalid operator!")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
func main() {
	calculate(10, 5, "+")
	calculate(10, 5, "-")
	calculate(10, 5, "*")
	calculate(10, 0, "/") // Division by zero example
}
