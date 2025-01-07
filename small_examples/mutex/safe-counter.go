package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	count int
	mutex sync.Mutex
}

func (sc *SafeCounter) Increment() {
	sc.mutex.Lock()
	sc.count++
	sc.mutex.Unlock()
}
func (sc *SafeCounter) GetValue() int {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	return sc.count
}
func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter.GetValue())
}
