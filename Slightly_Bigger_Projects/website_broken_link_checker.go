package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getLinks(baseURL string) ([]string, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch URL: %s, Status code: %d", baseURL, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
		}
	})
	return links, nil
}

func checkLink(link string, baseURL string) (string, int) {
	if !strings.HasPrefix(link, "http") {
		link = baseURL + link
	}

	resp, err := http.Head(link)
	if err != nil {
		return link, 0
	}
	defer resp.Body.Close()

	return link, resp.StatusCode
}

func main() {
	var baseURL string
	fmt.Print("Enter the website URL to check: ")
	fmt.Scan(&baseURL)

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalf("Invalid URL: %v", err)
	}

	links, err := getLinks(baseURL)
	if err != nil {
		log.Fatalf("Error fetching links: %v", err)
	}

	reportFile, _ := os.Create("broken_links_report.txt")
	defer reportFile.Close()

	fmt.Println("Checking links...")
	for _, link := range links {
		fullLink, statusCode := checkLink(link, parsedURL.Scheme+"://"+parsedURL.Host)
		if statusCode >= 400 || statusCode == 0 {
			log.Printf("Broken link: %s (Status: %d)\n", fullLink, statusCode)
			reportFile.WriteString(fmt.Sprintf("Broken link: %s (Status: %d)\n", fullLink, statusCode))
		}
	}
	fmt.Println("Broken links report saved to broken_links_report.txt")
}
