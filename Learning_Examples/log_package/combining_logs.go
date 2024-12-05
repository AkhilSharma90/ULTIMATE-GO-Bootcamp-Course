package main

import (
	"log"
	"runtime/debug"
)

func main() {
	log.Println("Program started")
	log.Println("Debugging information:")
	log.Println(string(debug.Stack()))
}

/*
2024/11/19 10:00:00 Program started
2024/11/19 10:00:00 Debugging information:
<stack trace details here>
*/
