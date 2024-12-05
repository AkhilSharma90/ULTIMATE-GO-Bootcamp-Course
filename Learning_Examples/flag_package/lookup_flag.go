package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.String("config", "default.conf", "Path to configuration file")

	// Parse flags
	flag.Parse()

	// Access flag details using Lookup
	configFlag := flag.Lookup("config")
	fmt.Println("Config Flag Value:", configFlag.Value)
	fmt.Println("Config Flag Usage:", configFlag.Usage)
}

//go run main.go -config=prod.conf
//Config Flag Value: prod.conf
//Config Flag Usage: Path to configuration file
