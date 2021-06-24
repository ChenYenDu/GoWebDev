package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sega struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sega{
		Name:  "Buddha",
		Motto: "The belief of no beliefs.",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
