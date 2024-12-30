package main

import (
	"fmt"
)

// Struct to represent a Project
type Project struct {
	Title string
	Tasks []string // Slice to hold tasks for the project
}

func main() {
	// Slice to store projects
	var projects []Project
	// Adding projects directly to the slice
	projects = append(projects, Project{
		Title: "Build Website",
		Tasks: []string{"Design Layout", "Develop Backend", "Deploy to Server"},
	})
	projects = append(projects, Project{
		Title: "Write Blog Post",
		Tasks: []string{"Research Topic", "Write Content", "Proofread"},
	})
	projects = append(projects, Project{
		Title: "Plan Event",
		Tasks: []string{"Book Venue", "Invite Guests", "Prepare Agenda"},
	})
	// Displaying all projects and their tasks
	fmt.Println("Project Management System")
	fmt.Println("--------------------------")
	for _, project := range projects {
		fmt.Printf("Project: %s\n", project.Title)
		for _, task := range project.Tasks {
			fmt.Printf("  - %s\n", task)
		}
	}
	// Adding a new task to the first project
	fmt.Println("\nAdding a new task to 'Build Website'")
	projects[0].Tasks = append(projects[0].Tasks, "Optimize Images")
	// Displaying updated tasks for the first project
	fmt.Printf("\nUpdated tasks for '%s':\n", projects[0].Title)
	for _, task := range projects[0].Tasks {
		fmt.Printf("  - %s\n", task)
	}
	// Filtering projects with more than 2 tasks
	fmt.Println("\nProjects with more than 2 tasks:")
	for _, project := range projects {
		if len(project.Tasks) > 2 {
			fmt.Printf("Project: %s, Tasks: %d\n", project.Title, len(project.Tasks))
		}
	}
}
