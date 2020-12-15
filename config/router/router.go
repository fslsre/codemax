package router

import (
	"codemax/config/routes"

	"github.com/gorilla/mux"
)

// NewRouter handle all routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.AllRoutes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.Handler)
	}
	return router
}
