// Install protobuf first and use protoc to generate the .pb.go file for this example.

package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

// Generated Protobuf file
type Person struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func main() {
	person := &Person{Name: "Alice", Age: 30}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("Marshaling error:", err)
	}

	fmt.Println("Serialized Protobuf:", data)
}
