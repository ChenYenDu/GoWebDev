package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user/goofy", set)
	http.HandleFunc("/user/goofy/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Println(c)
	fmt.Fprintln(w, "Your Cookie: ", c)
}

func set(w http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name:  "my-cookie",
		Value: "something",
	}
	http.SetCookie(w, c)
	fmt.Println(c)
	fmt.Fprintln(w, "Your Cookie: ", c)
	fmt.Fprintln(w, "Cookie Written - Check your browser")
	fmt.Fprintln(w, "in chrome go to: dev tools/ applications / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Your Cookie: ", c)
}
