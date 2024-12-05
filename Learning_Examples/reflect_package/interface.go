package main

import (
	"fmt"
	"reflect"
)

func Describe(i interface{}) {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
}

func main() {
	Describe(42)
	Describe("hello")
}
