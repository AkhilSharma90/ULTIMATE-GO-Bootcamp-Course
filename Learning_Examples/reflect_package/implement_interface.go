package main

import (
	"fmt"
	"reflect"
)

type Stringer interface {
	String() string
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func main() {
	p := Person{"Alice"}
	t := reflect.TypeOf(p)

	stringerType := reflect.TypeOf((*Stringer)(nil)).Elem()
	fmt.Println("Implements Stringer:", t.Implements(stringerType))
}
