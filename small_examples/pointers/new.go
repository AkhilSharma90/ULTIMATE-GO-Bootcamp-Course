package main

import "fmt"

func main() {
	ptr := new(int) // Allocate memory for an int
	*ptr = 100
	fmt.Println("Value at ptr:", *ptr) // Output: 100
}
