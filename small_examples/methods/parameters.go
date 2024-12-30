package main

import "fmt"

type Calculator struct{}

// Method to add two integers
func (c Calculator) Add(a, b int) int {
	return a + b
}
func main() {
	calc := Calculator{}
	fmt.Println("Sum:", calc.Add(10, 20)) // Output: 30
}
