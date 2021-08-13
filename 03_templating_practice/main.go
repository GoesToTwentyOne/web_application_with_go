package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type cardtitle struct {
	Name1 string
	Name2 string
	Name3 string
}

func main() {
	cardtest := cardtitle{
		Name1: "Test One",
		Name2: "Test Two",
		Name3: "Test Three",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", cardtest)
	if err != nil {
		log.Fatal(err)
	}

}
