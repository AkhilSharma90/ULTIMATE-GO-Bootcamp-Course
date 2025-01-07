package main

import "fmt"

type Config struct {
	Host string
	Port int
}

func main() {
	cfg := new(Config) // Allocate memory for a Config struct
	// Assign values
	cfg.Host = "localhost"
	cfg.Port = 8080
	fmt.Printf("Config: %+v\n", *cfg)
}
