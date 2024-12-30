package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1 // Base case
	}
	return n * factorial(n-1) // Recursive case
}
func main() {
	fmt.Println("Factorial of 5:", factorial(5)) // Output: 120
}
