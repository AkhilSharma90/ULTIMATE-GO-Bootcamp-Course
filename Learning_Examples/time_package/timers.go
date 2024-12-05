package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting...")
	<-timer.C
	fmt.Println("Timer expired!")
}

/*
Waiting...
Timer expired!
*/
