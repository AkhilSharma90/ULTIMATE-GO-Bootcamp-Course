package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 30}
	ptr := &p                      // Pointer to struct
	fmt.Println("Name:", ptr.Name) // Access fields directly
	fmt.Println("Age:", ptr.Age)
	ptr.Age = 31 // Modify struct fields via pointer
	fmt.Println("Updated Age:", p.Age)
}
