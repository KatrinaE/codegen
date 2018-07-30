package main

import (
	"os"
	"text/template"
)

//go:generate go run gen.go

var tpl = template.Must(
	template.New("").Parse(`Hello World!`))

func main() {
	data := struct{}{}
	out, _ := os.Create("output.txt")
	defer out.Close()
	tpl.Execute(out, data)
}
