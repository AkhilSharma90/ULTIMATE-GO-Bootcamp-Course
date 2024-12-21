package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Alice", Age: 30}

	// Serialize to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error serializing JSON:", err)
		return
	}

	fmt.Println("Serialized JSON:", string(jsonData))
}
