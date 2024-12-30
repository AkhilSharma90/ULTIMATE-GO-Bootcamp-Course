package main

import "fmt"

func main() {
	x := 10
	ptr := &x                         // Pointer to x
	*ptr = 20                         // Modify the value at the pointer
	fmt.Println("New value of x:", x) // Output: 20
}
