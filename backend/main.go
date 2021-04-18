package main

import (
	"log"
	"net/http"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/routes"
)

func main() {
	database.InitDatabase()

	router := routes.NewRouter()

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening on 8080...")
	s.ListenAndServe()
}