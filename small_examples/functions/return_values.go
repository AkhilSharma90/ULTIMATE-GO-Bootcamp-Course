package main

import "fmt"

// Function to calculate the square of a number
func square(num int) int {
	return num * num
}
func main() {
	result := square(5) // Call the function and store the result
	fmt.Println("Square:", result)
}
