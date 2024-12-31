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

// fibonacci(6) = fibonacci(5) + fibonacci(4)
// fibonacci(5) = fibonacci(4) + fibonacci(3)
// fibonacci(4) = fibonacci(3) + fibonacci(2)
// fibonacci(3) = fibonacci(2) + fibonacci(1)
// fibonacci(2) = fibonacci(1) + fibonacci(0)
// fibonacci(1) = 1, fibonacci(0) = 0
// fibonacci(2) = 1
// fibonacci(3) = 2
// fibonacci(4) = 3
// fibonacci(5) = 5
// fibonacci(6) = 5 + 3 = 8
