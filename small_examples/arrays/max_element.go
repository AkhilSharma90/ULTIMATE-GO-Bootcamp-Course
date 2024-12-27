package main

import "fmt"

func main() {
	numbers := [6]int{12, 45, 23, 56, 19, 8}
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	fmt.Println("Maximum value:", max)
}
