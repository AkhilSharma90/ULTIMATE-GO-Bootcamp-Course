package main

import (
	"fmt"
	"sync"
)

var (
	sum   int
	mutex sync.Mutex
)

func addValue(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	sum += value
	mutex.Unlock()
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go addValue(i, &wg)
	}
	wg.Wait()
	fmt.Println("Sum:", sum)
}
