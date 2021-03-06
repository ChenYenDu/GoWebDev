package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Template: think as a container which holding all templates whcih i pass
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
