package main

import "fmt"

type Task struct {
	Title     string
	Completed bool
}

func main() {
	tasks := []Task{
		{"Buy groceries", false},
		{"Clean the house", true},
		{"Pay bills", false},
	}
	fmt.Println("To-Do List:")
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("- %s [%s]\n", task.Title, status)
	}
}
