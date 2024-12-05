package main

import (
	"log"
)

func main() {
	log.Println("Program started")
	log.Fatal("A critical error occurred")
	log.Println("This will not be executed") // This line won't run
}

//2024/11/19 10:00:00 Program started
//2024/11/19 10:00:00 A critical error occurred
