package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("Unix Timestamp (Seconds):", now.Unix())
	fmt.Println("Unix Timestamp (Milliseconds):", now.UnixMilli())
	fmt.Println("Unix Timestamp (Nanoseconds):", now.UnixNano())
}

/*
Unix Timestamp (Seconds): 1732047165
Unix Timestamp (Milliseconds): 1732047165123
Unix Timestamp (Nanoseconds): 1732047165123456789
*/
