package main

type User struct {
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}

type Users []User

var users = Users{}
