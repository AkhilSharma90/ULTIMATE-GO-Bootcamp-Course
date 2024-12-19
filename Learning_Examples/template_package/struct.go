package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Define a template
	tmpl := `Name: {{.Name}}, Age: {{.Age}}`

	// Parse the template
	t, err := template.New("person").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Create a struct
	p := Person{Name: "Alice", Age: 30}

	// Execute the template
	err = t.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
