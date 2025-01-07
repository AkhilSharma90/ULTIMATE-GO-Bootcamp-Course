package main

import (
	"fmt"
	"sync"
)

func safeWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Worker %d recovered from panic: %v\n", id, r)
		}
	}()
	if id == 3 {
		panic("simulated panic")
	}
	fmt.Printf("Worker %d completed\n", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go safeWorker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers completed")
}
