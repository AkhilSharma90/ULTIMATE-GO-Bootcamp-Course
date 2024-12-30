package main

import (
	"fmt"
)

// Struct to represent a student
type Student struct {
	Name  string
	Marks int
}

func main() {
	// Slice to hold a list of students
	var students []Student
	// Adding students to the slice
	students = append(students, Student{Name: "Alice", Marks: 85})
	students = append(students, Student{Name: "Bob", Marks: 70})
	students = append(students, Student{Name: "Charlie", Marks: 92})
	students = append(students, Student{Name: "Diana", Marks: 60})
	// Display the list of students
	fmt.Println("List of Students:")
	for _, student := range students {
		fmt.Printf("Name: %s, Marks: %d\n", student.Name, student.Marks)
	}
	// Calculate the average marks
	var totalMarks int
	for _, student := range students {
		totalMarks += student.Marks
	}
	averageMarks := float64(totalMarks) / float64(len(students))
	fmt.Printf("\nAverage Marks: %.2f\n", averageMarks)
	// Categorize students based on their marks
	fmt.Println("\nStudent Categories:")
	for _, student := range students {
		fmt.Printf("Name: %s - ", student.Name)
		if student.Marks >= 85 {
			fmt.Println("Grade: A (Excellent)")
		} else if student.Marks >= 70 {
			fmt.Println("Grade: B (Good)")
		} else if student.Marks >= 50 {
			fmt.Println("Grade: C (Pass)")
		} else {
			fmt.Println("Grade: F (Fail)")
		}
	}
	// Find the top-scoring student
	var topStudent Student
	for _, student := range students {
		if student.Marks > topStudent.Marks {
			topStudent = student
		}
	}
	fmt.Printf("\nTop Scoring Student: %s with %d marks\n", topStudent.Name, topStudent.Marks)
}
