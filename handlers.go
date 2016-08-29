package main

import (
	"fmt"
	"html"
	"net/http"
)

//"encoding/json"

//"github.com/gorilla/mux"

func Create(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%q", html.EscapeString(req.URL.Path))
}
