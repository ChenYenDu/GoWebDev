# Step 1

Create ```func getUser``` and put it in a new file, session.go. This refactor allows us to use the same code in index and bar.

```Go
func getUser(w http.ResponseWriter, req *http.Request) user
```

# IMPORTANT

Now that you have two go files, you cannot use "go run main.go" t get your application to run. That is only asking for one go file: main.go. You need to use either "go build" and the run the executable, or "go run *.go"


# Step 2

Create ```func signup``` and removed the signup code from ```func index```. A new field for password was added to the user struct.

```Go
func signup(w http.ResponseWriter, req *http.Request)
```

# Step 3

Create ```func alreadyLoggedIn``` and put it on the session.go page. This refactor allows us to use the same code in bar and signup.

```Go
func alreadyLoggedIn(req *http.Request) bool
```