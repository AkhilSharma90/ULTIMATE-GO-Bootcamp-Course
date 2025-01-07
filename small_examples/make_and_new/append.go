package main

import "fmt"

func main() {
	nums := make([]int, 2, 5) // Length: 2, Capacity: 5
	nums[0] = 10
	nums[1] = 20
	nums = append(nums, 30) // Add an element
	fmt.Println(nums)       // Output: [10 20 30]
}
