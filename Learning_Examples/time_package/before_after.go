package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Date(2024, 11, 19, 14, 0, 0, 0, time.UTC)
	time2 := time.Date(2024, 11, 19, 16, 0, 0, 0, time.UTC)

	fmt.Println("Is time1 before time2?", time1.Before(time2))
	fmt.Println("Is time1 after time2?", time1.After(time2))
}

/*
Is time1 before time2? true
Is time1 after time2? false
*/
