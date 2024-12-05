package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	todoList := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List:")
		for i, item := range todoList {
			fmt.Printf("%d. %s\n", i+1, item)
		}

		fmt.Println("\nOptions:")
		fmt.Println("1. Add item")
		fmt.Println("2. Remove item")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Print("Enter a new item: ")
			scanner.Scan()
			newItem := strings.TrimSpace(scanner.Text())
			todoList = append(todoList, newItem)
		case "2":
			fmt.Print("Enter item number to remove: ")
			scanner.Scan()
			var itemNumber int
			fmt.Sscan(scanner.Text(), &itemNumber)
			if itemNumber > 0 && itemNumber <= len(todoList) {
				todoList = append(todoList[:itemNumber-1], todoList[itemNumber:]...)
			} else {
				fmt.Println("Invalid item number")
			}
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
