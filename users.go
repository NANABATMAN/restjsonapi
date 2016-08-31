package main

import "sync"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}

var (
	users = struct {
		sync.RWMutex
		m map[int]*User
	}{m: make(map[int]*User)}
	seq = 1
)
