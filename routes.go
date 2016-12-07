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

//var routes = Routes{
//	Route{
//		"Index",
//		"GET",
//		"/",
//		Index,
//	},
//	Route{
//		"TodoIndex",
//		"GET",
//		"/todos",
//		TodoIndex,
//	},
//	Route{
//		"TodoShow",
//		"GET",
//		"/todos/{todoId}",
//		TodoShow,
//	},
//	Route{
//		"TodoCreate",
//		"POST",
//		"/todos",
//		TodoCreate,
//	},
//}

var routes = Routes{
	Route{
		"GetUser",
		"GET",
		"/users/{userId}",
		GetUser,
	},
	Route{
		"UpdateUser",
		"GET",
		"/users",
		UpdateUser,
	},
	Route{
		"DeleteUser",
		"GET",
		"/",
		Index,
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		CreateUser,
	},
}