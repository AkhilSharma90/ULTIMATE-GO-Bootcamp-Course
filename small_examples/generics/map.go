package main

import "fmt"

// Map is a generic function that transforms a slice of type T to a slice of type U.
func Map[T any, U any](slice []T, transform func(T) U) []U {
	var result []U
	for _, v := range slice {
		result = append(result, transform(v))
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3}
	squared := Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println(squared) // Output: [1 4 9]
}
