package main

import (
	"net/http"

	"github.com/ChenYenDu/GoWebDev/42_mongodb/05_mongodb/01_update-user-controller/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	// create a new UseController
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	// added route
	r.POST("/user", uc.CreateUser)
	// added route plus parameter
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
