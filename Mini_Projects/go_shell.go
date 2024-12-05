package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to GoShell! Type 'exit' to quit.")
	fmt.Println("----------------------------------------")

	for {
		// Display the prompt
		fmt.Print("GoShell> ")

		// Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim whitespace and split command
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		args := strings.Split(input, " ")
		command := args[0]

		// Handle built-in commands
		switch command {
		case "exit":
			fmt.Println("Exiting GoShell. Goodbye!")
			return
		case "cd":
			if len(args) < 2 {
				fmt.Println("cd: missing argument")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("cd error:", err)
			}
			continue
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  cd <dir>  - Change directory")
			fmt.Println("  exit      - Exit the shell")
			fmt.Println("  help      - Display this help message")
			fmt.Println("  <command> - Execute any system command")
			continue
		}

		// Execute system commands
		cmd := exec.Command(command, args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Command error:", err)
		}
	}
}

//go run go_shell.go
//help, pwd, ls, echo hello - all these commands can be run
