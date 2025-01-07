package main

import "fmt"

func main() {
	inventory := make(map[string]int) // Map for inventory
	// Adding items
	inventory["Apples"] = 10
	inventory["Bananas"] = 5
	// Updating item quantity
	inventory["Apples"] += 5
	// Printing inventory
	fmt.Println("Inventory:")
	for item, quantity := range inventory {
		fmt.Printf("- %s: %d\n", item, quantity)
	}
}
