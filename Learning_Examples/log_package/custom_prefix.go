package main

import (
	"log"
)

func main() {
	log.SetPrefix("INFO: ")
	log.Println("This is an informational message")

	log.SetPrefix("ERROR: ")
	log.Println("This is an error message")
}

//output
//INFO: 2024/11/19 10:00:00 This is an informational message
//ERROR: 2024/11/19 10:00:00 This is an error message
