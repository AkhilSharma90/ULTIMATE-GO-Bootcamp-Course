package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}

	p.Age = 26     // Update a field
	fmt.Println(p) // Output: {Alice 26}
}
