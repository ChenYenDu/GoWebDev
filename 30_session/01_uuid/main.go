package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// For this code to run, you will need this package:
// github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true, // this means the cookie is only allow with https
			HttpOnly: true, // this means the cookie can not be access with JavaScript
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
