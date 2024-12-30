package main

import "fmt"

func main() {
	m := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	// Collect keys in a slice
	//transferring values from a map into a slice
	keys := []string{}

	//alice, bob, charlie
	for key := range m {
		keys = append(keys, key)
	}
	// Iterate over keys and access map values
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := m[key]
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
