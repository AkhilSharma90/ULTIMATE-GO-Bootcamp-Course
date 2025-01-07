package main

import "fmt"

type Box[T any] struct {
	Value T
}

func main() {
	intBox := Box[int]{Value: 100}
	stringBox := Box[string]{Value: "Go Generics"}
	fmt.Println(intBox)    // Output: {100}
	fmt.Println(stringBox) // Output: {Go Generics}
}
