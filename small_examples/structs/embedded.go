package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID string
}

func main() {
	e := Employee{
		Person:     Person{Name: "Alice", Age: 25},
		EmployeeID: "E123",
	}
	fmt.Println("Name:", e.Name) // Access embedded fields directly
}
