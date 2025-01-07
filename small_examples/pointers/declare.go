package main

import "fmt"

func main() {
	x := 10
	ptr := &x // Get the address of x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)
	fmt.Println("Value at ptr (dereference):", *ptr)
}


ptr actually points to x
by storing its address

when you print out the value of ptr, you will get the address of x
but when you dereference ptr, by putting a star, you will get the VALUE of x