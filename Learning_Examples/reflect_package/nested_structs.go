package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City string
}

type User struct {
	Name    string
	Address Address
}

func main() {
	u := User{Name: "Alice", Address: Address{City: "New York"}}
	v := reflect.ValueOf(&u).Elem()

	// Access nested field
	cityField := v.FieldByName("Address").FieldByName("City")
	if cityField.CanSet() {
		cityField.SetString("San Francisco")
	}

	fmt.Println("Updated Struct:", u)
}
