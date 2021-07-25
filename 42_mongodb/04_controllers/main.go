package main

import (
	"net/http"

	"github.com/ChenYenDu/GoWebDev/42_mongodb/04_controllers/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	// create a new UseController
	uc := controllers.NewUserController()

	r.GET("/user/:id", uc.GetUser)
	// added route
	r.POST("/user", uc.CreateUser)
	// added route plus parameter
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

// func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	s := `<!DOCTYPE html>
// 	<html lang="en">
// 	<head>
// 		<meta charset="UTF-8">
// 		<title>Index</title>
// 	</head>
// 	<body>
// 		<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
// 	</body>
// 	</html>
// 	`

// 	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(s))
// }
