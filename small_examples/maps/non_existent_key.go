package main

import "fmt"

func main() {
	m := map[string]string{}
	fmt.Println(m["nonexistent"]) // Outputs: 0 (default value for int)
}
