package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ErrorResponse should be used for any non-successful response
type ErrorResponse struct {
	Code    int
	Message string
}

// NewRouter creates a router with all server paths
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	// See if this works
	r.HandleFunc("/song", AuthMiddleware(CreateSongHandler)).Methods("POST")
	r.HandleFunc("/user", CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/login", LoginUserHandler).Methods("POST")

	return r
}
