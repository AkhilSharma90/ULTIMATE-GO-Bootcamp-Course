package main

import "fmt"

func main() {
	x := 10
	ptr := &x // Get the address of x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)
	fmt.Println("Value at ptr (dereference):", *ptr)
}
