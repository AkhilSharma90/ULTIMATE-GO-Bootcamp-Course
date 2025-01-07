package main

import (
	"fmt"
	"sync"
)

var (
	mutexA sync.Mutex
	mutexB sync.Mutex
)

func taskA() {
	mutexA.Lock()
	fmt.Println("Task A: Locked A")
	mutexB.Lock()
	fmt.Println("Task A: Locked B")
	mutexB.Unlock()
	mutexA.Unlock()
}
func taskB() {
	mutexB.Lock()
	fmt.Println("Task B: Locked B")
	mutexA.Lock()
	fmt.Println("Task B: Locked A")
	mutexA.Unlock()
	mutexB.Unlock()
}
func main() {
	go taskA()
	go taskB()
	select {} // Prevent the program from exiting
}
