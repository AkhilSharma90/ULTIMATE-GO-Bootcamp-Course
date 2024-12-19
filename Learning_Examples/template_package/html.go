package main

import (
	"html/template"
	"os"
)

func main() {
	// Define a template
	tmpl := `<h1>Hello, {{.}}</h1>`

	t, err := template.New("htmlExample").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Execute the template with untrusted input
	err = t.Execute(os.Stdout, "<script>alert('XSS')</script>")
	if err != nil {
		panic(err)
	}
}
