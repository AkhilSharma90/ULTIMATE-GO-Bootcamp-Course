package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var people []Person // Empty slice
	// Adding elements using append
	people = append(people, Person{"Alice", 25})
	people = append(people, Person{"Bob", 30})
	fmt.Println(people)
}
