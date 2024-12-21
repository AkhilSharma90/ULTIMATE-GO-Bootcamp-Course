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
	jsonData := `{"name": "Alice", "age": 30}`

	var person Person
	// Deserialize JSON
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}

	fmt.Printf("Deserialized Struct: %+v\n", person)
}
