package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Running a command
	cmd := exec.Command("echo", "Hello from Go!")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	fmt.Println("Command executed successfully.")
}
