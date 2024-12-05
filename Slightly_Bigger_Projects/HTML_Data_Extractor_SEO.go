package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func extractMetadata(url string) (string, string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", "", fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", "", "", fmt.Errorf("error parsing HTML: %w", err)
	}

	var title, description, keywords string

	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			title = node.FirstChild.Data
		}

		if node.Type == html.ElementNode && node.Data == "meta" {
			for _, attr := range node.Attr {
				if attr.Key == "name" && attr.Val == "description" {
					for _, a := range node.Attr {
						if a.Key == "content" {
							description = a.Val
						}
					}
				}
				if attr.Key == "name" && attr.Val == "keywords" {
					for _, a := range node.Attr {
						if a.Key == "content" {
							keywords = a.Val
						}
					}
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	return title, description, keywords, nil
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
	}

	file, err := os.Create("metadata.csv")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	_ = writer.Write([]string{"URL", "Title", "Description", "Keywords"})

	for _, url := range urls {
		title, description, keywords, err := extractMetadata(url)
		if err != nil {
			log.Printf("Error extracting metadata from %s: %v", url, err)
			continue
		}

		_ = writer.Write([]string{url, title, description, keywords})
	}

	fmt.Println("Metadata extraction complete. Results saved to metadata.csv")
}
