package main

import "fmt"

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Sum of numbers:", sum)
}
