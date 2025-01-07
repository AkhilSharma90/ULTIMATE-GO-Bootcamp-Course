package main

import "fmt"

func main() {
	nums := make([]int, 3, 5) // Length: 3, Capacity: 5
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", nums, len(nums), cap(nums))
}
