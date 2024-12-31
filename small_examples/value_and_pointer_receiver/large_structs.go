package main

import "fmt"

type LargeStruct struct {
	Data [1000]int
}

// Method with pointer receiver
func (ls *LargeStruct) Reset() {
	for i := range ls.Data {
		ls.Data[i] = 0
	}
}
func main() {
	large := LargeStruct{}
	large.Reset() // Efficient as it avoids copying the struct
	fmt.Println("First element after reset:", large.Data[0])
}
