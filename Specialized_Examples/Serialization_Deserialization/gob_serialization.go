package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alice", Age: 30}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	// Serialize to Gob
	err := encoder.Encode(person)
	if err != nil {
		fmt.Println("Error serializing Gob:", err)
		return
	}

	fmt.Println("Serialized Gob:", buffer.Bytes())
}
