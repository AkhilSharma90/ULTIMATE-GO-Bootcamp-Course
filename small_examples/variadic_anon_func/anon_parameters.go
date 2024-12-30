package main

import "fmt"

func main() {
	// Anonymous function with parameters
	func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}("Alice")
}
