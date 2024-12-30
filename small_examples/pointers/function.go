package main

import "fmt"

func increment(num *int) {
	*num++ // Increment the value at the pointer
}
func main() {
	x := 10
	fmt.Println("Before:", x)
	increment(&x)
	fmt.Println("After:", x) // Output: 11
}
