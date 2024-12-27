package main

import "fmt"

func main() {
	original := [3]int{10, 20, 30}
	copy := original // Creates a copy
	copy[0] = 99
	fmt.Println("Original:", original)
	fmt.Println("Copy:", copy)
}
