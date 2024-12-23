package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <json_file>")
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Println("Invalid JSON:", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(prettyJSON))
}
