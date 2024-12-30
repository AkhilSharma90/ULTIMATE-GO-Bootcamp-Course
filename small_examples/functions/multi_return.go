package main

import "fmt"

// Function to return two numbers
func divideAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}
func main() {
	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}
