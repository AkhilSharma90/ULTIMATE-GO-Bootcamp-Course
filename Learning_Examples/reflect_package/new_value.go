package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(42)
	newValue := reflect.New(t).Elem()
	newValue.SetInt(100)

	fmt.Println("New Value:", newValue.Int())
}
