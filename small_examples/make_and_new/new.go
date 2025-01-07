package main

import "fmt"

func main() {
	num := new(int) // Allocate memory for an integer
	fmt.Println("Default value:", *num)
	*num = 42 // Assign a value through the pointer
	fmt.Println("Updated value:", *num)
}
