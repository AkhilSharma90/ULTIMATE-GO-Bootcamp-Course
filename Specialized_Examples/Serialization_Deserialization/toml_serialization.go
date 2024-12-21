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
	config := Config{Database: "example_db", Port: 5432}

	data, err := toml.Marshal(config)
	if err != nil {
		fmt.Println("Error serializing TOML:", err)
		return
	}

	fmt.Println("Serialized TOML:\n", string(data))
}
