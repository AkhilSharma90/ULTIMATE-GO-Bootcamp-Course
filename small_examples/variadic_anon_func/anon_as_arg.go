package main

import "fmt"

func executeOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	result := executeOperation(5, 3, func(x, y int) int {
		return x * y
	})
	fmt.Println("Multiplication result:", result) // Output: 15
}
