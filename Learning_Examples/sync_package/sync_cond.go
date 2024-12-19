package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var ready bool

	go func() {
		time.Sleep(2 * time.Second) // Simulate work
		cond.L.Lock()
		ready = true
		cond.L.Unlock()
		cond.Signal() // Notify waiting goroutine
	}()

	cond.L.Lock()
	for !ready {
		cond.Wait() // Wait until signaled
	}
	cond.L.Unlock()
	fmt.Println("Goroutine is ready!")
}

//sync.Cond allows goroutines to wait for a condition and be notified when the condition changes.