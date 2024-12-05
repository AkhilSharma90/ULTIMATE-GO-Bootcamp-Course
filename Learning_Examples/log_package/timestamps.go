package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("Log with date, time, and microseconds")
}

//2024/11/19 10:00:00.123456 Log with date, time, and microseconds
