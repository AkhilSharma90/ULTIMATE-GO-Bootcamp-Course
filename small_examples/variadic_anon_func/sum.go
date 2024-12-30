package main

import "fmt"

// Variadic function to calculate the sum
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3)) // Output: 6
	fmt.Println("Sum of 10, 20:", sum(10, 20))   // Output: 30
	fmt.Println("Sum of nothing:", sum())        // Output: 0
}
