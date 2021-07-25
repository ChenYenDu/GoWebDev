# Curl

## GET

```bash
curl http://localhost:8080/user/1
```

## POST

```bash
curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond", "Gender":"male", "Age":32, "Id":"777"}' http://localhost:8080/user
```

-X is short for --request Specifies a custom request method to use when communicating with the HTTP server.

-H is short for --header

-d is short for --data

## DELETE

```bash
curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
```

IMPORTANT: Make sure you update your import statements to import packages from the correct location!
