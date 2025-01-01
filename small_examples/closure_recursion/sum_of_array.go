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

// sum(nums, 5) → nums[4] (5) + sum(nums, 4)
// sum(nums, 4) → nums[3] (4) + sum(nums, 3)
// sum(nums, 3) → nums[2] (3) + sum(nums, 2)
// sum(nums, 2) → nums[1] (2) + sum(nums, 1)
// sum(nums, 1) → nums[0] (1) + sum(nums, 0)

// sum(nums, 1) = 1
// sum(nums, 2) = 2 + 1 = 3
// sum(nums, 3) = 3 + 3 = 6
// sum(nums, 4) = 4 + 6 = 10
// sum(nums, 5) = 5 + 10 = 15
