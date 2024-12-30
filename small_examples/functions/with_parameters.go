package main

import "fmt"

// Function with one parameter
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
func main() {
	greet("Alice") // Call the function with an argument
}
