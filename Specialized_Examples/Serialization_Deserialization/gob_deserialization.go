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
	decoder := gob.NewDecoder(&buffer)

	// Serialize to Gob
	encoder.Encode(person)

	var decodedPerson Person
	// Deserialize Gob
	err := decoder.Decode(&decodedPerson)
	if err != nil {
		fmt.Println("Error deserializing Gob:", err)
		return
	}

	fmt.Printf("Deserialized Struct: %+v\n", decodedPerson)
}
