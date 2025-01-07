package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}
func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}
func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Counter Value:", counter.GetValue())
}
