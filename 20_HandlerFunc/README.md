# HandlerFunc

[http.HandlerFunc](https://pkg.go.dev/net/http#HandlerFunc)

```Go
type HandlerFunc func(ResponseWriter, *Request)
```

```Go
func (f Handler) ServeHTTP(w ResponseWriter, r *Request)
```

This is just a nice thing to know about. You wouldn't do this in production code probably.

---

## Question

Could you get http.Handle to take a func with this signature: func(ResponseWriter, *Request)?
