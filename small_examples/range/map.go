package main

import "fmt"

func main() {
	m := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
