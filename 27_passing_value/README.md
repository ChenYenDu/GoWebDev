# Passing values

## Form submission

We can pass values from the client to server through URL or through the body of the request.

When you submit a form, you can use either the "GET" or "POST" method. The "POST" method sends the form submission through the body of request. The "GET" method for a form submission sends the form submission values through the url.

I remember that like this:

```markdown
post
body

get
url
```

__POST__ has four letters and so does __form__.
__GET__ has three letters and so does __url__.

---

## URL value

You can always append values to a URL.

Anything after the ```?``` is the query string - the area where the values are stored.

The values are stord in a identifier=value fashion.

You can have multiple identifier=value by separating them with the & ampersand.

---

## Retrieving values

While there are multiple ways to retrieve values, we will stick with:

[func (*Request) FormValue](https://godoc.org/net/http#Request.FormValue)

```Go
func (r *Request) FormValue(key string) string
```

FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.
