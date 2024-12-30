package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method to create a new person with default age
func NewPerson(name string) Person {
	return Person{Name: name, Age: 18}
}
func main() {
	p := NewPerson("Alice")
	fmt.Printf("New Person: %+v\n", p) // Output: {Name:Alice Age:18}
}
