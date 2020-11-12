package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger logs request infos on terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate check if user is authenticated to user an restrict route
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
