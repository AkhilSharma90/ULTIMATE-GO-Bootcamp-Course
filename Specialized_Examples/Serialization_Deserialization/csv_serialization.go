package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	file, err := os.Create("people.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Name", "Age"})

	// Write records
	for _, person := range people {
		writer.Write([]string{person.Name, fmt.Sprint(person.Age)})
	}

	fmt.Println("CSV file created successfully")
}
