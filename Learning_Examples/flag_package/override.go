package main

import (
	"flag"
	"fmt"
)

func main() {
	// Create a custom flag set
	customFlags := flag.NewFlagSet("Custom", flag.ExitOnError)

	// Define flags in the custom flag set
	port := customFlags.Int("port", 8080, "Port number")
	verbose := customFlags.Bool("verbose", false, "Enable verbose output")

	// Parse the flags
	customFlags.Parse([]string{"-port=9090", "-verbose"})

	fmt.Println("Port:", *port)
	fmt.Println("Verbose:", *verbose)
}

//Port: 9090
//Verbose: true
