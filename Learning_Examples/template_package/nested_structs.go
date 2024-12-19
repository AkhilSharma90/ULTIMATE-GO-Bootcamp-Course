package main

import (
	"os"
	"text/template"
)

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Address Address
}

func main() {
	// Define a template
	tmpl := `Name: {{.Name}}, Location: {{.Address.City}}, {{.Address.Country}}`

	t, err := template.New("nested").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Execute the template
	person := Person{Name: "John", Address: Address{City: "New York", Country: "USA"}}
	err = t.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}
