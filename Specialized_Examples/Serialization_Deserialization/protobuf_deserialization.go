package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)


type Person struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func main() {
	data := []byte{...} // Serialized Protobuf data

	var person Person
	err := proto.Unmarshal(data, &person)
	if err != nil {
		log.Fatal("Unmarshaling error:", err)
	}

	fmt.Printf("Deserialized Struct: %+v\n", person)
}
