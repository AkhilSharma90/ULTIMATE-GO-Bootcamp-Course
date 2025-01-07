package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	students := make([]Student, 0) // Dynamic slice of students
	// Adding students
	students = append(students, Student{"Alice", 95})
	students = append(students, Student{"Bob", 85})
	// Printing all students
	for i := 0; i < len(students); i++ {
		fmt.Printf("Student %d: Name: %s, Score: %d\n", i+1, students[i].Name, students[i].Score)
	}
}
