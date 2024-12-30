package main

import "fmt"

func createPointer() *int {
	x := 42
	return &x // Return address of x
}
func main() {
	ptr := createPointer()
	fmt.Println("Value at pointer:", *ptr) // Output: 42
}
