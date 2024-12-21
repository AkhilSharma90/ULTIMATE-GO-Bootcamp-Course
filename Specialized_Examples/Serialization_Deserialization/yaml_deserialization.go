package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func main() {
	yamlData := `
name: Alice
age: 30
`

	var person Person
	// Deserialize YAML
	err := yaml.Unmarshal([]byte(yamlData), &person)
	if err != nil {
		fmt.Println("Error deserializing YAML:", err)
		return
	}

	fmt.Printf("Deserialized Struct: %+v\n", person)
}
