package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	xmlData := `<person>
  <name>Alice</name>
  <age>30</age>
</person>`

	var person Person
	// Deserialize XML
	err := xml.Unmarshal([]byte(xmlData), &person)
	if err != nil {
		fmt.Println("Error deserializing XML:", err)
		return
	}

	fmt.Printf("Deserialized Struct: %+v\n", person)
}
