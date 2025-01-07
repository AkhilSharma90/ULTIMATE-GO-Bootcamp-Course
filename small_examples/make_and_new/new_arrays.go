package main

import "fmt"

func main() {
	arr := new([3]int) // Allocate memory for an array of 3 integers
	fmt.Println("Default array:", *arr)
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("Updated array:", *arr)
}
