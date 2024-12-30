package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	var akhil Person
	akhil = Person{Name: "Akhil", Age: 25}
	//p := Person{Name: "Alice", Age: 25} // Initialize with field values
	fmt.Println(akhil) // Output: {Alice 25}
}
