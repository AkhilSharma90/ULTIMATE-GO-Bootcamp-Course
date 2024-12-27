package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25} // Initialize with field values
	fmt.Println(p)                      // Output: {Alice 25}
}
