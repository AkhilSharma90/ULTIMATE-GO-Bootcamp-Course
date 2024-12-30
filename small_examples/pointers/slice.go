package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	ptr := &slice                           // Pointer to slice
	fmt.Println("Slice via pointer:", *ptr) // Output: [1 2 3]
	(*ptr)[1] = 42                          // Modify second element
	fmt.Println("Updated slice:", slice)    // Output: [1 42 3]
}
