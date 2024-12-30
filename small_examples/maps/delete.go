package main

import "fmt"

func main() {
	m := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	delete(m, "Bob") // Remove the key "Bob"
	fmt.Println(m)   // Output: map[Alice:25]
}
