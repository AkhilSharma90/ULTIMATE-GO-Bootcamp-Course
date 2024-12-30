package main

import "fmt"

// Recursive function to calculate Fibonacci numbers
func fibonacci(n int) int {
	if n <= 1 {
		return n // Base case
	}
	return fibonacci(n-1) + fibonacci(n-2) // Recursive case
}
func main() {
	fmt.Println("Fibonacci of 6:", fibonacci(6)) // Output: 8
}
