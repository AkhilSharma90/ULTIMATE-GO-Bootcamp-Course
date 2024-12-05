package main

import (
	"fmt"
	"reflect"
)

func Greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	funcValue := reflect.ValueOf(Greet)

	// Invoke the function with arguments
	args := []reflect.Value{reflect.ValueOf("Alice")}
	funcValue.Call(args)
}
