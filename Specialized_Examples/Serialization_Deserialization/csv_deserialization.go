package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	file, err := os.Open("people.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header row
	_, _ = reader.Read()

	var people []Person

	// Read records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		age, _ := strconv.Atoi(record[1])
		people = append(people, Person{Name: record[0], Age: age})
	}

	fmt.Printf("Deserialized Structs: %+v\n", people)
}
