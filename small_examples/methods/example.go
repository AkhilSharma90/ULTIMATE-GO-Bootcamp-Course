package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method for the Person type
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s, and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	alice := Person{Name: "Alice", Age: 30}
	alice.Greet() // Call the method

	bob := Person{Name: "Bob", Age: 20}
	bob.Greet()
}
