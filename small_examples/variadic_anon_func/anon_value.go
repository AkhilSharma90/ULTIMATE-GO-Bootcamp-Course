package main

import "fmt"

func main() {
	result := func(a, b int) int {
		return a + b
	}(5, 3)
	fmt.Println("Sum:", result) // Output: 8
}
