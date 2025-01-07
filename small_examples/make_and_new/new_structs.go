package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := new(Person) // Allocate memory for a Person struct
	fmt.Println("Default struct:", *p)
	*p.Name = "Alice" // Access fields via the pointer
	p.Age = 30
	fmt.Println("Updated struct:", *p)
}
