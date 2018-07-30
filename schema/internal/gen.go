package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"text/template"
)

//go:generate go run gen.go
//go:generate go fmt codegen/user.gen.go

func main() {
	var resp map[string]interface{}
	in, _ := os.Open("../apidata/instagram/user.json")
	b, _ := ioutil.ReadAll(in)
	json.Unmarshal(b, &resp)

	data := struct {
		Name   string
		Fields map[string]interface{}
	}{
		"User",
		resp,
	}

	tpl, _ := template.New("template.tpl").Funcs(template.FuncMap{
		"Title": strings.Title,
		"TypeOf": func(v interface{}) string {
			if v == nil {
				return "string"
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
	}).ParseFiles("template.tpl")

	out, _ := os.Create("codegen/user.gen.go")
	defer out.Close()

	tpl.Execute(out, data)
}
