package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	// Initialize the notes struct
	note := notes{
		To:      "Akhil",
		From:    "Joe",
		Heading: "Meeting",
		Body:    "Lets hang out at 8!",
	}

	// Marshal the struct to XML format with indentation
	file, err := xml.MarshalIndent(note, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	// Write the XML to a file
	err = os.WriteFile("notes1.xml", file, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File successfully written to notes1.xml")
}
