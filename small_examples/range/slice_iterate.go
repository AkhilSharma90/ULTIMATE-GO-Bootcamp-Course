package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
}
