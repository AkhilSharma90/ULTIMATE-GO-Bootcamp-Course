package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock() // Ensure the mutex is unlocked after the critical section
	fmt.Println("Critical section")
}
