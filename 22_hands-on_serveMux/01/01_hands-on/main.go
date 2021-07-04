package main

import (
	"fmt"
	"net/http"
)

func n(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "this is index page.")
}

func d(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "dog dog dog")
}

func m(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "mememee~")
}

func main() {
	http.HandleFunc("/", n)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
