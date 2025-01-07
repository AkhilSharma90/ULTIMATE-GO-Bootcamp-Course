package main

import (
	"fmt"
	"time"
)

func service(name string, duration time.Duration) {
	fmt.Printf("%s starting\n", name)
	time.Sleep(duration)
	fmt.Printf("%s completed\n", name)
}
func main() {
	go service("Service A", 2*time.Second)
	go service("Service B", 1*time.Second)
	go service("Service C", 3*time.Second)
	fmt.Println("All services initiated")
	time.Sleep(4 * time.Second)
}
