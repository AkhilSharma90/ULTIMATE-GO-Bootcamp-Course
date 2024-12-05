package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const (
	cgroupMemoryPath = "/sys/fs/cgroup/memory/mycontainer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <command> [args...]")
		fmt.Println("Example: go run container.go run /bin/bash")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runContainer()
	case "child":
		runChildProcess()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func runContainer() {
	fmt.Println("Starting container...")

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
	}

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error starting container: %v", err)
	}
}

func runChildProcess() {
	fmt.Println("Inside the container!")
	setupCgroups()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Change the hostname
	if err := syscall.Sethostname([]byte("mycontainer")); err != nil {
		log.Fatalf("Error setting hostname: %v", err)
	}

	// Mount a new filesystem
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		log.Fatalf("Error mounting proc: %v", err)
	}
	defer syscall.Unmount("/proc", 0)

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}

func setupCgroups() {
	fmt.Println("Setting up cgroups...")

	// Create the cgroup directory
	err := os.MkdirAll(cgroupMemoryPath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Error creating cgroup directory: %v", err)
	}

	// Set memory limit to 64MB
	err = os.WriteFile(filepath.Join(cgroupMemoryPath, "memory.limit_in_bytes"), []byte("67108864"), 0700)
	if err != nil {
		log.Fatalf("Error setting memory limit: %v", err)
	}

	// Add the current process to the cgroup
	err = os.WriteFile(filepath.Join(cgroupMemoryPath, "tasks"), []byte(fmt.Sprintf("%d", os.Getpid())), 0700)
	if err != nil {
		log.Fatalf("Error adding process to cgroup: %v", err)
	}
}

/*
build the project - go build -o container container.go
start the container with bash running inside it - sudo ./container run /bin/bash

now you can run multiple commands inside the container -
hostname
ps aux
ls /proc
cat /sys/fs/cgroup/memory/mycontainer/memory.limit_in_bytes

you can exit the container by typing exit

you can run this command to directly run command inside the container - sudo ./container run echo "Hello from the container!"
*/
