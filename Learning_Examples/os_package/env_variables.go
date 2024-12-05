package main

import (
	"fmt"
	"os"
)

func main() {
	// Setting an environment variable
	os.Setenv("MY_ENV_VAR", "12345")

	// Getting an environment variable
	value := os.Getenv("MY_ENV_VAR")
	fmt.Println("MY_ENV_VAR:", value)

	// Listing all environment variables
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
