package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func fileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func findDuplicates(root string) {
	hashes := make(map[string][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		hash, err := fileHash(path)
		if err != nil {
			fmt.Printf("Error hashing file %s: %v\n", path, err)
			return nil
		}

		hashes[hash] = append(hashes[hash], path)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}

	for hash, files := range hashes {
		if len(files) > 1 {
			fmt.Printf("Duplicate Files for Hash [%s]:\n", hash)
			for _, file := range files {
				fmt.Println(file)
			}
			fmt.Println()
		}
	}
}

func main() {
	rootDir := "./" // Change this to the directory you want to scan
	fmt.Printf("Scanning for duplicates in directory: %s\n", rootDir)
	findDuplicates(rootDir)
}
