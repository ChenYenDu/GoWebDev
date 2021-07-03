# HandleFunc

[http.HandleFunc](https://pkg.go.dev/net/http#HandleFunc)

```Go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

Takes a func with parameters type: ```http.ResponseWriter``` and ```*http.Request```.

