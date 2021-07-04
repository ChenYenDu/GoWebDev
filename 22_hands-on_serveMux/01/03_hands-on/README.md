Take the previous program and change it so that:

- func main uses http.Handle instead of http.HandleFunc

Contstraint: Do not change anything outside the func main

Hints:

[http.HandlerFunc](https://pkg.go.dev/net/http#HandlerFunc)

```Go
type HandlerFunc func(ResponseWriter, *Request)
```

[http.HandleFunc](https://pkg.go.dev/net/http#HandleFunc)

```Go
func HandleFunc(pattern string, handler func(ResponseWrite, *Request))
```

[source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)

```Go
func (mux *ServeMux) HandleFunc(pattern string, handler func(ReqponseWriter, *Request)) {
    mux.Handle(pattern, HandlerFunc(handler))
}
```