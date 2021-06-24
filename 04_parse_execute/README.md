# Type Template

## template.Template

```golang
template.Template
```

---

## Parsing templates

### template.ParseFiles

```Go
func ParseFiles(filenames ...string) (*Template, error)
```

### template.ParseGlob

```Go
func ParseGlob(pattern string) (*Template, error)
```

---

### template.Parse

```Go
func (t *Template) Parse(text string) (*Template, error)
```

### template.ParseFiles

```Go
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

### template.ParseGlob

```Go
func (t *Template) ParseGlob(pattern string) (*Template, error)
```


---

## Executing templates

### template.Execute

```Go
func (t *Template) Execute(wr io.Writer, data interface{}) err
```

### template.ExecuteTemplate

```Go
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

---

## Helpful template functions

### template.Must

```Go
func Must(t *Template, err error) *Template
```

### template.New

```Go
func New(name string) *Template
```

---

## The init function

```Go
func init()
```
