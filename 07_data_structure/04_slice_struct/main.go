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

	gandhi := sega{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sega{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sega{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sega{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	segas := []sega{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, segas)
	if err != nil {
		log.Fatalln(err)
	}
}
