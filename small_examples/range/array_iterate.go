package main

import "fmt"

func main() {
	arr := [3]string{"Go", "Python", "Java"}
	for i, lang := range arr {
		fmt.Printf("Index: %d, Language: %s\n", i, lang)
	}
}
