package main

import (
	"fmt"
	"sync"
	"time"
)

func longTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)
	time.Sleep(3 * time.Second)
	fmt.Printf("Task %d finished\n", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go longTask(i, &wg)
	}
	wg.Wait()
	fmt.Println("All long tasks completed")
}
