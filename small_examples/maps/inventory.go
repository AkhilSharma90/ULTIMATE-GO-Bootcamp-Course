package main

import "fmt"

func main() {
	inventory := map[string]int{
		"Apples":  10,
		"Bananas": 5,
	}
	// Add items
	inventory["Oranges"] = 7
	// Update items
	inventory["Apples"] += 5
	// Collect keys in a slice
	keys := []string{}
	for key := range inventory {
		keys = append(keys, key)
	}
	// Iterate over keys and print inventory
	fmt.Println("Inventory:")
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		fmt.Printf("%s: %d\n", key, inventory[key])
	}
}
