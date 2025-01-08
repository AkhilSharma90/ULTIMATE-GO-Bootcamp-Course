package main

import "fmt"

func Identity[T any](value T) T {
	return value
}
func main() {
	fmt.Println(Identity(42))          // Output: 42
	fmt.Println(Identity("Hello, Go")) // Output: Hello, Go
}
