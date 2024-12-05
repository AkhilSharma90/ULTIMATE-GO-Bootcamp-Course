package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func calculateDirSize(path string) (int64, error) {
	var totalSize int64 = 0

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return totalSize, nil
}

func main() {
	directory := "./" // Change this to any directory path

	fmt.Printf("Calculating size of directory: %s\n", directory)
	size, err := calculateDirSize(directory)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Directory Size: %.2f MB\n", float64(size)/1024/1024)
}
