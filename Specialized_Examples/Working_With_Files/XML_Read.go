package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	// Read the file content
	data, err := os.ReadFile("notes.xml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Initialize an empty Notes struct
	note := &Notes{}

	// Unmarshal the XML data into the Notes struct
	err = xml.Unmarshal(data, note)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	// Print the note fields
	fmt.Println("To:", note.To)
	fmt.Println("From:", note.From)
	fmt.Println("Heading:", note.Heading)
	fmt.Println("Body:", note.Body)
}
