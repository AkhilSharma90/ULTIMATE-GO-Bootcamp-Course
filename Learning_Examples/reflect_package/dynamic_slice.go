package main

import (
	"fmt"
	"reflect"
)

func main() {
	sliceType := reflect.SliceOf(reflect.TypeOf(0)) // Slice of int
	slice := reflect.MakeSlice(sliceType, 0, 5)

	for i := 0; i < 5; i++ {
		slice = reflect.Append(slice, reflect.ValueOf(i*10))
	}

	fmt.Println("Slice:", slice.Interface())
}
