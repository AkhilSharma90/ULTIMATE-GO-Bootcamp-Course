package main

import (
	"log"
)

func main() {
	log.Println("This is a basic log message")
	log.Printf("Logging with formatted output: %s", "Hello, Go!")
}

//output below
//2024/11/19 10:00:00 This is a basic log message
//2024/11/19 10:00:00 Logging with formatted output: Hello, Go!
