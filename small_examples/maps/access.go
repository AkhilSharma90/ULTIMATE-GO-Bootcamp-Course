package main

import "fmt"

func main() {
	m := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("Alice's age:", m["Alice"]) // Output: Alice's age: 25
}
