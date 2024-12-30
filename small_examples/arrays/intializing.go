package main

import "fmt"

//first example shows how to initialize, next examples show how to do it in one line

// func main() {
// 	// Declare an array of integers
// 	var nums [5]int
// 	nums[0] = 10
// 	nums[1] = 20
// 	nums[2] = 30
// 	nums[3] = 40
// 	nums[4] = 50

// 	fmt.Println("Array:", nums)

// }

func main() {
	// var nums [5]int = [5]int{10, 20, 30, 40, 50}
	// fmt.Println("Initialized array:", nums)

	names := [4]string{"Alice", "Bob", "Charlie"}
	fmt.Println("Names:", names)
}
