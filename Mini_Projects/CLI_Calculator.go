package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var operator string

	fmt.Println("Enter first number:")
	fmt.Scanln(&a)
	fmt.Println("Enter second number:")
	fmt.Scanln(&b)
	fmt.Println("Enter operator (+, -, *, /):")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", a+b)
	case "-":
		fmt.Printf("Result: %.2f\n", a-b)
	case "*":
		fmt.Printf("Result: %.2f\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Result: %.2f\n", a/b)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
}
