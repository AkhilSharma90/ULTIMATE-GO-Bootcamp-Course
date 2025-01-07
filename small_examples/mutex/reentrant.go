package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}
func (c *Counter) IncrementTwice() {
	c.Increment()
	c.Increment()
}
func main() {
	counter := Counter{}
	counter.IncrementTwice()
	fmt.Println("Counter Value:", counter.count)
}
