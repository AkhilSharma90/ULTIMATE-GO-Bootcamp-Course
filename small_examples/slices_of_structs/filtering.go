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
	fmt.Println("Incomplete Tasks:")
	for _, task := range tasks {
		if !task.Completed {
			fmt.Printf("- %s\n", task.Title)
		}
	}
}
