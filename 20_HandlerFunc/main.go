package main

import (
	"io"
	"net/http"
)

// This methods have the same underlying type: func(http.ResponseWriter, *http,Request)
func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat cat cat")
}

func main() {

	// This is just type converting
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
