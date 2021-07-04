# Serving files

## ServeContent

[http.ServeContent](https://pkg.go.dev/net/http?utm_source=godoc#ServeContent)

```Go
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
```

---

## ServeFile

[http.ServeFile](https://godoc.org/net/http#ServeFile)

```Go
func ServeFile(w ResponseWriter, r *Request, name string)
```

---

## FileServer & StripPrefix

[http.FileServer](https://godoc.org/net/http#FileServer)

```Go
func FileServer(root FileSystem) Handler
```

[http.StripPrefix](https://godoc.org/net/http#StripPrefix)

```Go
func StripPrefix(prefix string, h Handler) Handler
```
