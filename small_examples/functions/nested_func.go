package main

import "fmt"

func factorial(n int) int {

	var innerFactorial func(int) int

	innerFactorial = func(x int) int {
		if x == 0 {
			return 1
		}
		return x * innerFactorial(x-1)
	}
	return innerFactorial(n)
}

func main() {

	fmt.Println("Factorial of 5:", factorial(5)) // Output: 120
}
