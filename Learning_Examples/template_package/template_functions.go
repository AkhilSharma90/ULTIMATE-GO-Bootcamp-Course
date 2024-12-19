package main

import (
	"os"
	"text/template"
	"strings"
)

func main() {
	// Create a template with a function
	funcMap := template.FuncMap{
		"toUpper": strings.ToUpper,
	}

	tmpl := `Original: {{.}}, Uppercase: {{. | toUpper}}`
	t, err := template.New("funcExample").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Execute the template
	err = t.Execute(os.Stdout, "hello world")
	if err != nil {
		panic(err)
	}
}
