package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person   // Declare a struct variable
	fmt.Println(p) // Default values: {"" 0}
}
