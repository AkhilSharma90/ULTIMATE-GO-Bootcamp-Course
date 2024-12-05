package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("custom.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	customLogger := log.New(file, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)
	customLogger.Println("This is a custom log message")
	customLogger.Printf("Custom log with format: %s", "Hello, custom logger!")
}

/*
CUSTOM: 2024/11/19 10:00:00 main.go:18: This is a custom log message
CUSTOM: 2024/11/19 10:00:00 main.go:19: Custom log with format: Hello, custom logger!
*/
