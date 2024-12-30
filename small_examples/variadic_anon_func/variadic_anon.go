package main

import "fmt"

func main() {
	sum := func(numbers ...int) int {
		total := 0
		for _, num := range numbers {
			total += num
		}
		return total
	}
	fmt.Println("Sum:", sum(1, 2, 3, 4)) // Output: 10
}
