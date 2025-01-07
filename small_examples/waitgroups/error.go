package main

import (
	"fmt"
	"sync"
)

func task(id int, wg sync.WaitGroup, errors []error, mu *sync.Mutex) {
	defer wg.Done()
	if id%2 == 0 {
		mu.Lock()
		errors = append(errors, fmt.Errorf("task %d failed", id))
		mu.Unlock()
	} else {
		fmt.Printf("Task %d completed successfully\n", id)
	}
}
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, &wg, &errors, &mu)
	}
	wg.Wait()
	fmt.Println("All tasks finished")
	for _, err := range errors {
		fmt.Println("Error:", err)
	}
}
