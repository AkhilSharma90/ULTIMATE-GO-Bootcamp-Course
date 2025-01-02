package main

import "fmt"

type Config struct {
	Port    int
	Enabled bool
}

// Method to set default values
func (c *Config) SetDefaults() {
	c.Port = 8080
	c.Enabled = true
}

func main() {
	cfg := Config{}
	cfg.SetDefaults()
	fmt.Println("Config after defaults:", cfg) // Output: {8080 true}
}
