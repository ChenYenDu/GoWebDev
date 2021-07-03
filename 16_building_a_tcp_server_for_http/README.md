# HTTP Server

HTTP uses TCP.

To create a server that works with HTTP, we just create a TCP server.

To configure our code to handle request/response in an HTTP fashion which works with browsers, we need to adhere to HTTP standards.

## HTTP/1.1 message

An HTTP message is made up of the following:

- A request/status line
- Zero or more header fields
- An empty line indicating the end of the header section
- an optional message body

---

### Request line (request)

GET / HTTP / 1.1

[method SP request-target SP HTTP-version CRLF](https://datatracker.ietf.org/doc/html/rfc7230#section-3.1.1)

### Status line (response)

HTTP / 1.1 302 Found

[HTTP-version SP status-code SP reson-phrase CRLF](https://datatracker.ietf.org/doc/html/rfc7230#section-3.1.2)

---

## Writing a response

```Go
body := "CHECK OUT THE RESPONSE BODY PAYLOAD" 
io.WriteString(conn, "HTTP/1.1 200 OK\r\n")             // status line
fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))  // header
fmt.Fprintf(conn, "Content-Type: text/plain\r\n")       // header
io.WriteString(conn, "\r\n")    // blank line; CRLP; carriage-return line-feed
io.WriteString(conn, body)      // body, aka, payload
```

---

## Useful for parsing the request-line & status-line

### Parsing String

[strings.Fields](https://pkg.go.dev/strings#Fields)

```Go
func Fields(s string) []string
```
