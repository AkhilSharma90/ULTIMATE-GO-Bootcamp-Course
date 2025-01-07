package main

import (
	"fmt"
	"sync"
)

type TicketSystem struct {
	tickets int
	mutex   sync.Mutex
}

func (ts *TicketSystem) BookTicket() bool {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	if ts.tickets > 0 {
		ts.tickets--
		return true
	}
	return false
}
func main() {
	ts := TicketSystem{tickets: 5}
	var wg sync.WaitGroup
	for i := 1; i <= 7; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if ts.BookTicket() {
				fmt.Printf("User %d booked a ticket\n", id)
			} else {
				fmt.Printf("User %d could not book a ticket\n", id)
			}
		}(i)
	}
	wg.Wait()
}
