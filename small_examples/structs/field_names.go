package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Bob", 30} // Without field names
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
