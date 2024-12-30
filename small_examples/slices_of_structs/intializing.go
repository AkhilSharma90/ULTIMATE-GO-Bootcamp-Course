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
		{"Charlie", 35},
	}
	fmt.Println(people)
}
