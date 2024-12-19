package main

import (
	"os"
	"text/template"
)

type Item struct {
	Name     string
	Quantity int
}

func main() {
	// Define a template with an `if` condition
	tmpl := `{{if gt .Quantity 0}}Item: {{.Name}}, Quantity: {{.Quantity}}{{else}}Item: {{.Name}}, Out of Stock{{end}}`

	t, err := template.New("inventory").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Execute the template
	t.Execute(os.Stdout, Item{Name: "Apple", Quantity: 5})
	t.Execute(os.Stdout, Item{Name: "Banana", Quantity: 0})
}
