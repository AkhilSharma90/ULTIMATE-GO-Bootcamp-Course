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
	people[1].Age = 35 // Update Bob's age
	fmt.Println(people)
}
