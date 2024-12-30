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
	for i := 0; i < len(people); i++ {
		fmt.Printf("Person %d: %s, Age: %d\n", i+1, people[i].Name, people[i].Age)
	}
}
