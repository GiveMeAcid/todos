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
		"/users/{Id}",
		GetUser,
	},
	Route{
		"GetPeople",
		"GET",
		"/users",
		GetPeople,
	},
	//Route{
	//	"DeleteUser",
	//	"DELETE",
	//	"/users/{id}",
	//	DeleteUser,
	//},
	Route{
		"CreateUser",
		"POST",
		"/users/{Id}",
		CreateUser,
	},
}
