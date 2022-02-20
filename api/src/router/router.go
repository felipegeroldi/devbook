package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter generate new router with
// configured routes
func GenerateRouter() *mux.Router {
	router := mux.NewRouter()
	routes.ConfigureRoutes(router)

	return router
}
