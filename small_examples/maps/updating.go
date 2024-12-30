package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Alice"] = 25 // Add a key-value pair
	m["Bob"] = 30   // Add another pair
	fmt.Println(m)  // Output: map[Alice:25 Bob:30]
}
