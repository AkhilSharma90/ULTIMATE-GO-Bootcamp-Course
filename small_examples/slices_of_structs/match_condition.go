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
	fmt.Println("People older than 30:")
	for _, person := range people {
		if person.Age > 30 {
			fmt.Printf("%s, Age: %d\n", person.Name, person.Age)
		}
	}
}
