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
	person := Person{Name: "Alice", Age: 30}

	// Serialize to XML
	xmlData, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error serializing XML:", err)
		return
	}

	fmt.Println("Serialized XML:\n", string(xmlData))
}
