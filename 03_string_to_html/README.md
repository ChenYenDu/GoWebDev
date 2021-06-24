# Techniques we learn here

---

* concatenate
* CLI pipeline - output to a file with `>`

## Cod we will use from the standard library

---

### os.Create

This allow us to create a file.

```golang
func Create(name string) (*File, error)
```

### defer

The defer keyword allows us to defer the execution of a statement until the function in which we have placed the defer statement returns.

### io.Copy

This allows us to copy from a source to a destination.

```golang
func Copy(dst Writer, src Reader) (written int64, err error)
```

### strings.NewReader

NewReader returns a new Reader reading from s.

```golang
func NewReader(s string) *Reader
```

### os.Args

Args is a variable from package os. Args hold the command-line arguments, starting with the program name.

```golang
var Args []string
```
