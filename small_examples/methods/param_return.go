package main

import "fmt"

type Calculator struct{}

// Method to multiply two integers
func (c Calculator) Multiply(a, b int) int {
	return a * b
}
func main() {
	calc := Calculator{}
	fmt.Println("Multiplication result:", calc.Multiply(3, 4)) // Output: 12
}
