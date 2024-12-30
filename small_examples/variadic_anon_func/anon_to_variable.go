package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Addition:", add(10, 20)) // Output: 30
}
