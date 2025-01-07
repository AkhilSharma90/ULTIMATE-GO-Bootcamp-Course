package main

import "fmt"

func main() {
	slice := []int{10, 20}
	ptr := &slice                        // Pointer to slice
	*ptr = append(*ptr, 30)              // Append an element via pointer
	fmt.Println("Updated slice:", slice) // Output: [10 20 30]
}
