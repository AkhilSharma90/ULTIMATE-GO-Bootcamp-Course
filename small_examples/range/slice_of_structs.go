package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	students := []Student{
		{"Alice", 90},
		{"Bob", 80},
		{"Charlie", 85},
	}
	for _, student := range students {
		fmt.Printf("Name: %s, Score: %d\n", student.Name, student.Score)
	}
}
