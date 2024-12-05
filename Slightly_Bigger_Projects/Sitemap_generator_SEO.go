package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type URL struct {
	Loc string `xml:"loc"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func crawl(baseURL string, visited map[string]bool, maxDepth int, currentDepth int) []string {
	if currentDepth > maxDepth {
		return nil
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", baseURL, err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Non-200 status code for URL %s: %d\n", baseURL, resp.StatusCode)
		return nil
	}

	links := extractLinks(resp.Body, baseURL)
	var allLinks []string

	for _, link := range links {
		if !visited[link] {
			visited[link] = true
			allLinks = append(allLinks, link)
			allLinks = append(allLinks, crawl(link, visited, maxDepth, currentDepth+1)...)
		}
	}

	return allLinks
}

func extractLinks(body io.Reader, baseURL string) []string {
	doc, err := html.Parse(body)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return nil
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					u, err := url.Parse(attr.Val)
					if err != nil || u.Scheme == "" || u.Host == "" {
						continue
					}
					if strings.HasPrefix(u.Host, baseURL) {
						links = append(links, u.String())
					}
				}
			}
		
