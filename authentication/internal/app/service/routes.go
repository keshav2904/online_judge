package service

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
		"Login",
		"POST",
		"/login",
		HandleLogin,
	},
	Route{
		"Registration",
		"POST",
		"/register",
		HandleRegistration,
	},
	// Route{
	// 	"ValidateToken",
	// 	"POST",
	// 	"/validateToken",
	// 	HandleValidateToken,
	// },
	Route{
		"AddBatch",
		"POST",
		"/addbatch",
		isAuthorized(isTeacher(HandleAddBatch)),
	},
	Route{
		"GetBatches",
		"GET",
		"/getbatch",
		isAuthorized(HandleGetBatches),
	},
	Route{
		"AddAssignment",
		"POST",
		"/addassignment/{bID}",
		isAuthorized(isTeacher(HandleAddAssignment)),
	},
}
