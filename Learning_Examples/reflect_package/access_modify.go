package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 30}
	v := reflect.ValueOf(&p).Elem() // Use Elem() to access the value inside the pointer

	nameField := v.FieldByName("Name")
	ageField := v.FieldByName("Age")

	fmt.Println("Name:", nameField.String())
	fmt.Println("Age:", ageField.Int())

	// Modifying fields
	if nameField.CanSet() {
		nameField.SetString("Bob")
	}
	if ageField.CanSet() {
		ageField.SetInt(35)
	}

	fmt.Println("Modified Struct:", p)
}
