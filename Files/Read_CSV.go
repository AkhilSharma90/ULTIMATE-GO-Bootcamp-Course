package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the CSV file
	csvfile, err := os.Open("test.csv")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer csvfile.Close()

	// Create a new CSV reader
	csvreader := csv.NewReader(csvfile)

	// Read all the records from the CSV
	rows, err := csvreader.ReadAll()
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	// Print each row of the CSV
	for _, row := range rows {
		fmt.Printf("Name: %s, City: %s, Language: %s\n", row[0], row[1], row[2])
	}
}
