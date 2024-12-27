package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	book := Book{Title: "Go Programming", Author: "John Doe", Pages: 300}
	fmt.Printf("Book: %s by %s (%d pages)\n", book.Title, book.Author, book.Pages)
}
