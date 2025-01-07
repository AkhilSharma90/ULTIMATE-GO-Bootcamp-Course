package main

import (
	"fmt"
	"sync"
)

func subTask(id, subID int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d - Subtask %d started\n", id, subID)
}
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	var subWG sync.WaitGroup
	for i := 1; i <= 2; i++ {
		subWG.Add(1)
		go subTask(id, i, &subWG)
	}
	subWG.Wait()
	fmt.Printf("Worker %d finished\n", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers and subtasks completed")
}
