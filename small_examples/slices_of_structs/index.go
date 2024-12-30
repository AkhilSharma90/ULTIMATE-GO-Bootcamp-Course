package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	fmt.Println("First person:", people[0].Name) // Access the first person's name
}
