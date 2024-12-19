package main

import (
	"os"
	"text/template"
)

func main() {
	// Define a template with `range`
	tmpl := `Shopping List:
{{range .}}- {{.}}
{{end}}`

	t, err := template.New("list").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	items := []string{"Milk", "Eggs", "Bread"}

	// Execute the template
	err = t.Execute(os.Stdout, items)
	if err != nil {
		panic(err)
	}
}
