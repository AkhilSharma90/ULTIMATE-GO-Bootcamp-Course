package main

import "fmt"

// Recursive function to calculate sum of array elements
func sum(arr []int, n int) int {
	if n == 0 {
		return 0 // Base case
	}
	return arr[n-1] + sum(arr, n-1) // Recursive case
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of array:", sum(nums, len(nums))) // Output: 15
}
