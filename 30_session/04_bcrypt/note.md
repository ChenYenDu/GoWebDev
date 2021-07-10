You should never store a password without enrypting it.

We will use ```bcrypt``` to encrypt our passwords.

# Step 1:
You will need to get this package:

```bash
go get golang.org/x/crpyto/bcrypt
```

# Step 2:

Change your user struct's password field to the data type []byte

# Step 3:

Use bcrypt to encrypt your password before storing it.
