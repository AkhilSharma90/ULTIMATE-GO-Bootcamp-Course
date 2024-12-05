package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x *int
	var y int

	fmt.Println("x is nil:", reflect.ValueOf(x).IsNil())
	fmt.Println("x is pointer:", reflect.ValueOf(x).Kind() == reflect.Ptr)
	fmt.Println("y is pointer:", reflect.ValueOf(y).Kind() == reflect.Ptr)
}
