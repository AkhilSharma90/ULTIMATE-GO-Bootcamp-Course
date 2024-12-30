package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	students := []Student{
		{"Alice", 85},
		{"Bob", 92},
		{"Charlie", 78},
	}
	// Add bonus points to each student's score
	for i := range students {
		students[i].Score += 5
	}
	fmt.Println("Updated Student Scores:")
	for _, student := range students {
		fmt.Printf("%s: %d\n", student.Name, student.Score)
	}
}
