package main

import (
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	log.Println("Program started")
	log.Panic("A panic-inducing error occurred")
	log.Println("This will not be executed")
}

/*
2024/11/19 10:00:00 Program started
2024/11/19 10:00:00 A panic-inducing error occurred
2024/11/19 10:00:00 Recovered from panic: A panic-inducing error occurred]
*/
