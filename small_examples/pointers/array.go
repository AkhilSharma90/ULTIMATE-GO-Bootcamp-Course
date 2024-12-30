package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	ptr := &arr                             // Pointer to the array
	fmt.Println("Array via pointer:", *ptr) // Output: [10 20 30]
	(*ptr)[1] = 99                          // Modify second element
	fmt.Println("Updated array:", arr)      // Output: [10 99 30]
}
