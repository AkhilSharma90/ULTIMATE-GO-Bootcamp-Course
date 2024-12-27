package main

import "fmt"

func main() {
	nums := [3]int{10, 20, 30}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, nums[i])
	}
}
