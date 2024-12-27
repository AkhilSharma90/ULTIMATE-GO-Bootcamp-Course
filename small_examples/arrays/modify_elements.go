package main

import "fmt"

func main() {
	nums := [3]int{10, 20, 30}
	nums[1] = 50 // Modify the second element
	fmt.Println("Modified array:", nums)
}
