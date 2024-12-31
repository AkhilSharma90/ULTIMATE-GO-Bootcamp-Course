package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Sum of slice:", sum(nums...)) // Output: 10
}
