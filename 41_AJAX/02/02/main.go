package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// execute index.gohtml template
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `Here are message from foo`
	fmt.Fprintln(w, s)
}
