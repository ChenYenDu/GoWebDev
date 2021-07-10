package main

import (
	"fmt"
	"net/http"
	"time"
)

func getUser(req *http.Request) user {
	var u user

	// get cookie
	c, err := req.Cookie("session")

	if err != nil {
		return u
	}

	// if the user exists already, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un.un]
	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")

	if err != nil {
		return false
	}

	un := dbSessions[c.Value]

	_, ok := dbUsers[un.un]

	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purpose
	showSessions()              // for demonstration purpose
	for k, v := range dbSessions {
		if time.Since(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}

		dbSessionsCleaned = time.Now()
		fmt.Println("AFTER CLEAN") // for demonstration purpose
		showSessions()             // for demonstration purpose
	}
}

// for demonstraction purposes
func showSessions() {
	fmt.Println("**********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
