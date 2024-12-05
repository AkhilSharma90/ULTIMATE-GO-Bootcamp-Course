package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	u := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, Type: %s, Tag: %s\n", field.Name, field.Type, field.Tag.Get("json"))
	}
}
