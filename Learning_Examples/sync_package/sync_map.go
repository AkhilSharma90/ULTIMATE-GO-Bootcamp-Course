package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Storing key-value pairs
	m.Store("Alice", 25)
	m.Store("Bob", 30)

	// Loading a value
	if age, ok := m.Load("Alice"); ok {
		fmt.Println("Alice's Age:", age)
	}

	// Iterating over all key-value pairs
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%s is %d years old\n", key, value)
		return true
	})
}

//sync.Map is a concurrent map designed for use cases where multiple goroutines need to read and write without locking.