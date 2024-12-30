package main

import "fmt"

func main() {
	// Define a map with string keys and integer values
	prices := map[string]int{
		"apple":  100,
		"banana": 50,
		"orange": 80,
	}

	// Print the original map
	fmt.Println("Original Prices:", prices)

	// Update the price of an "apple"
	prices["apple"] = 120

	// Update the price of a "banana"
	prices["banana"] = 60

	// Print the updated map
	fmt.Println("Updated Prices:", prices)
}
