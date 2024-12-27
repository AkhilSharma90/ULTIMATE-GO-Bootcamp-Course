package main

import "fmt"

type Address struct {
	City    string
	ZipCode int
}
type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  25,
		Address: Address{
			City:    "New York",
			ZipCode: 10001,
		},
	}
	fmt.Println("City:", p.Address.City)
}
