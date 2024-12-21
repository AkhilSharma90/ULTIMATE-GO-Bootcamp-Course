package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Person struct {
	Age  int32
	Name string
}

func main() {
	data := []byte{30, 0, 0, 0, 'A', 'l', 'i', 'c', 'e'}

	buffer := bytes.NewBuffer(data)

	var person Person

	// Deserialize Age
	binary.Read(buffer, binary.LittleEndian, &person.Age)

	// Deserialize Name
	person.Name = string(buffer.Bytes())

	fmt.Printf("Deserialized Struct: %+v\n", person)
}
