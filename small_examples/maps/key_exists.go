package main

import "fmt"

func main() {
	m := map[string]int{
		"Alice": 25,
	}
	age, exists := m["Bob"]
	if exists {
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob not found")
	}
}
