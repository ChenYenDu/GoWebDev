package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	// type convert: string -> integer
	count, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Fatalln("Checkout cookie, it's not an integer")

	}

	count++

	// type convert: integer -> string
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)

}
