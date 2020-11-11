package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents API routes
type Route struct {
	URI                string
	Method             string
	Function           func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

// ConfigRoutes configures all routes
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
