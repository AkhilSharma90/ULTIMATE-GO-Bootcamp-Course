package main

import "fmt"

func main() {
	nums := make([]int, 0, 5) // Length: 0, Capacity: 5
	for i := 1; i <= 5; i++ {
		nums = append(nums, i) // Dynamically grow the slice
		fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", nums, len(nums), cap(nums))
	}
}
