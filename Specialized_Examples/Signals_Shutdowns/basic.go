package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel to receive OS signals
	signalChan := make(chan os.Signal, 1)

	// Notify the channel on SIGINT and SIGTERM
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Waiting for signals...")
	sig := <-signalChan // Block until a signal is received
	fmt.Printf("Received signal: %s\n", sig)
}
