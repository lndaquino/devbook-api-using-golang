package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate return a new router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
