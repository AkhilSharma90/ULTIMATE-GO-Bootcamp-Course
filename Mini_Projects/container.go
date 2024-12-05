package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run container.go run <command>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runContainer()
	case "child":
		runChild()
	default:
		panic("Invalid command")
	}
}

func runContainer() {
	fmt.Printf("Running %v as a container\n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Clone a new namespace for the container
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // UTS namespace (hostname, domain)
			syscall.CLONE_NEWPID | // PID namespace
			syscall.CLONE_NEWNS | // Mount namespace
			syscall.CLONE_NEWNET, // Network namespace
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func runChild() {
	fmt.Printf("Running child process %v\n", os.Args[2:])

	// Change root filesystem (requires a Linux environment and a prepared root filesystem)
	if err := syscall.Chroot("/path/to/rootfs"); err != nil {
		fmt.Println("Error changing root:", err)
		os.Exit(1)
	}

	if err := os.Chdir("/"); err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	// Set hostname
	if err := syscall.Sethostname([]byte("container")); err != nil {
		fmt.Println("Error setting hostname:", err)
		os.Exit(1)
	}

	// Execute the provided command
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
}
