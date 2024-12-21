package main

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Database string `toml:"database"`
	Port     int    `toml:"port"`
}

func main() {
	tomlData := `
database = "example_db"
port = 5432
`

	var config Config
	err := toml.Unmarshal([]byte(tomlData), &config)
	if err != nil {
		fmt.Println("Error deserializing TOML:", err)
		return
	}

	fmt.Printf("Deserialized Struct: %+v\n", config)
}
