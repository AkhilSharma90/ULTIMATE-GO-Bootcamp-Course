package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method with a pointer receiver to modify the fields
func (p *Person) Birthday() {
	p.Age++ // Increment the age
}
func main() {
	person := Person{Name: "Alice", Age: 25}
	person.Birthday()                      // Modifies the original struct
	fmt.Println("After birthday:", person) // Output: {Alice 26}
}
