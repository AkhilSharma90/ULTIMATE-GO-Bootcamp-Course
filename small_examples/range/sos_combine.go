package main

import "fmt"

type Task struct {
	ID    int
	Title string
}

func main() {
	team1Tasks := []Task{
		{1, "Design Database Schema"},
		{2, "Set Up CI/CD Pipeline"},
	}
	team2Tasks := []Task{
		{3, "Develop API Endpoints"},
		{4, "Create Frontend Components"},
	}
	// Combine tasks
	allTasks := append(team1Tasks, team2Tasks...)
	fmt.Println("All Tasks:")
	for _, task := range allTasks {
		fmt.Printf("Task #%d: %s\n", task.ID, task.Title)
	}
}
