package main

import "fmt"

// A function that takes another function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Function to add two numbers
func add(a, b int) int {
	return a + b
}

func main() {
	result := applyOperation(5, 3, add)
	fmt.Println("Result:", result)
}
