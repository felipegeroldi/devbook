package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes represents an API route
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// ConfigureRoutes put all routes on the router
func ConfigureRoutes(router *mux.Router) {
	routes := userRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
}
