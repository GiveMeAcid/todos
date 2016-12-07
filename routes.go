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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{Id}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"GetUser",
		"GET",
		"/users/{login}",
		GetUser,
	},
	Route{
		"GetUsers",
		"GET",
		"/users",
		GetUsers,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/users/{login}",
		DeleteUser,
	},
	Route{
		"CreateUser",
		"POST",
		"/users/{login}",
		CreateUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/users/{login}",
		UpdateUser,
	},
}
