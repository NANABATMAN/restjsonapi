package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const (
	READ   = "/read/"
	DELETE = "/delete/"
)

func Create(res http.ResponseWriter, req *http.Request) {
	var userStruct User
	if decoder(res, req, &userStruct) != nil {
		log.Println("Error! Cannot decode JSON!")
		res.WriteHeader(500)
	}
	users = append(users, userStruct)
}

func Read(res http.ResponseWriter, req *http.Request) {
	username := getUsername(req.URL.Path, READ)
	for _, user := range users {
		if user.Username == username {
			res.Header().Set("Content-Type", "application/json; charset=utf-8")
			encoder := json.NewEncoder(res)
			err := encoder.Encode(user)
			if err != nil {
				log.Println("Error! Cannot encode JSON!")
				res.WriteHeader(500)
			}
		}
	}
}

func Update(res http.ResponseWriter, req *http.Request) {
	var userStruct User
	if decoder(res, req, &userStruct) != nil {
		log.Println("Error! Cannot decode JSON!")
		res.WriteHeader(500)
	}

	for i, user := range users {
		if user.Username == userStruct.Username {
			users[i] = userStruct
		}
	}
}

func Delete(res http.ResponseWriter, req *http.Request) {
	username := getUsername(req.URL.Path, DELETE)
	var deleteIndex int
	for i, user := range users {
		if user.Username == username {
			deleteIndex = i
		}
	}
	users = append(users[:deleteIndex], users[deleteIndex+1:]...)
}

func decoder(res http.ResponseWriter, req *http.Request, userStruct *User) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(userStruct)
	return err
}

func getUsername(path string, pattern string) string {
	return strings.Split(path, pattern)[1]
}
