package main

import "fmt"

type Task struct {
	Title       string
	Description string
	Completed   bool
}

func main() {
	tasks := []Task{
		{Title: "Learn Go", Description: "Complete Go tutorials", Completed: true},
		{Title: "Build a project", Description: "Work on a Go project", Completed: false},
	}
	for i, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("Task %d: %s - %s [%s]\n", i+1, task.Title, task.Description, status)
	}
}
