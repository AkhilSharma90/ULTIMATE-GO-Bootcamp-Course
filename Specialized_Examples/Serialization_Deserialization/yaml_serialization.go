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
	person := Person{Name: "Alice", Age: 30}

	// Serialize to YAML
	yamlData, err := yaml.Marshal(person)
	if err != nil {
		fmt.Println("Error serializing YAML:", err)
		return
	}

	fmt.Println("Serialized YAML:\n", string(yamlData))
}
