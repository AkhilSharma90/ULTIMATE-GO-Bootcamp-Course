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
	person := Person{Age: 30, Name: "Alice"}

	var buffer bytes.Buffer

	// Serialize Age
	binary.Write(&buffer, binary.LittleEndian, person.Age)

	// Serialize Name
	buffer.WriteString(person.Name)

	fmt.Println("Serialized Binary:", buffer.Bytes())
}
