package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	data := []byte{...} // Serialized Protobuf data

	var person Person
	err := proto.Unmarshal(data, &person)
	if err != nil {
		log.Fatal("Unmarshaling error:", err)
	}

	fmt.Printf("Deserialized Struct: %+v\n", person)
}
