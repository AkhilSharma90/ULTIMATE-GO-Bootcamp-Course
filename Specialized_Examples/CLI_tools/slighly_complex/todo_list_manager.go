package main

import (
	"fmt"
	"os"
)

var todoList []string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <add|list|delete> [item]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add <item>")
			return
		}
		item := os.Args[2]
		todoList = append(todoList, item)
		fmt.Printf("Added: %s\n", item)

	case "list":
		fmt.Println("To-Do List:")
		for i, item := range todoList {
			fmt.Printf("%d. %s\n", i+1, item)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <index>")
			return
		}
		index := os.Args[2]
		fmt.Printf("Deleted item %s (Not implemented!)\n", index) // Exercise left
	default:
		fmt.Println("Unknown command:", command)
	}
}
