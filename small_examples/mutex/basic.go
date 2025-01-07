package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
