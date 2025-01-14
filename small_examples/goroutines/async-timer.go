package main

import (
	"fmt"
	"time"
)

func delayedPrint(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message)
}
func main() {
	go delayedPrint("This prints after 1 second", 1*time.Second)
	go delayedPrint("This prints after 2 seconds", 2*time.Second)
	fmt.Println("Main function continues immediately")
	// time.Sleep(3 * time.Second)
}
