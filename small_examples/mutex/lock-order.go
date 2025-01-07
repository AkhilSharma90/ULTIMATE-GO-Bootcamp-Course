package main

import (
	"fmt"
	"sync"
)

var (
	mutexA sync.Mutex
	mutexB sync.Mutex
)

func task(order string) {
	mutexA.Lock()
	fmt.Println(order, "Locked A")
	mutexB.Lock()
	fmt.Println(order, "Locked B")
	mutexB.Unlock()
	mutexA.Unlock()
}
func main() {
	go task("Task 1")
	go task("Task 2")
	select {}
}
