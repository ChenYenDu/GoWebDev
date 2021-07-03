package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// This made hotdog a Handler
// Handler:
// type Handler interface {
// 	ServeHTTP(http.Response, *http.Request)
// }

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func.")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
