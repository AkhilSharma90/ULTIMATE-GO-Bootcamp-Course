package main

import (
	"os"
	"text/template"
)

func main() {
	// Define a template
	tmpl, err := template.New("hello").Parse("Hello, {{.}}!")
	if err != nil {
		panic(err)
	}

	// Execute the template with a data string
	err = tmpl.Execute(os.Stdout, "World")
	if err != nil {
		panic(err)
	}
}
