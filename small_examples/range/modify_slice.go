package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	for i := range nums {
		nums[i] *= 2 // Double each value
	}
	fmt.Println("Modified slice:", nums)
}
