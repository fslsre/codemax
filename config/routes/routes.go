package routes

import (
	"codemax/controllers/emailcontroller"
	"net/http"
)

// Route type
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// Routes type
type Routes []Route

// AllRoutes all routes
var AllRoutes = Routes{
	Route{
		Name:    "home",
		Method:  "GET",
		Pattern: "/",
		Handler: emailcontroller.Get,
	},
}
