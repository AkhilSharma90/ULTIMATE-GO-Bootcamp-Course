package main

import "fmt"

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}
func main() {
	sum := add(3, 4)
	fmt.Println("Sum:", sum)
}
