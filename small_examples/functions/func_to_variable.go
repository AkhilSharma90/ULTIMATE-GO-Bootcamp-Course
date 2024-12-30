package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	operation := add                        // Assign the function to a variable
	fmt.Println("Result:", operation(5, 3)) // Call the function via the variable
}
