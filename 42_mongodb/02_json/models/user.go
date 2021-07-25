package models

type User struct {
	Name   string `json:"Name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Id     string `json:"id"`
}
