package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Create",
		"POST",
		"/create",
		Create,
	},
	Route{
		"Read",
		"GET",
		"/read/{username}",
		Read,
	},
	Route{
		"Update",
		"PUT",
		"/update",
		Update,
	},
	Route{
		"Delete",
		"DELETE",
		"/delete/{username}",
		Delete,
	},
}
