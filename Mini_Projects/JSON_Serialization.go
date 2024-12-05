package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	// Encode to JSON
	user := User{Name: "John Doe", Email: "john@example.com", Age: 30}
	jsonData, _ := json.Marshal(user)
	fmt.Println("JSON Output:", string(jsonData))

	// Decode from JSON
	jsonInput := `{"name":"Jane Smith","email":"jane@example.com","age":25}`
	var decodedUser User
	json.Unmarshal([]byte(jsonInput), &decodedUser)
	fmt.Printf("Decoded User: %+v\n", decodedUser)
}
