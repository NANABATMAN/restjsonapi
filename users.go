package main

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}

var (
	users = map[int]*User{}
	seq   = 1
)
