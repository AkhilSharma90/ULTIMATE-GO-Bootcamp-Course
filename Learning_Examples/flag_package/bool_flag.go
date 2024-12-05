package main

import (
	"flag"
	"fmt"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")

	flag.Parse()

	if *debug {
		fmt.Println("Debug mode is ON")
	} else {
		fmt.Println("Debug mode is OFF")
	}
}

//go run main.go -debug
//Debug mode is ON
