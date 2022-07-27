package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"usebottles.com/steamgrid-proxy/config"
	"usebottles.com/steamgrid-proxy/controller"
)

func main() {
	cnf := *config.Cnf
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/search/{gameName}", controller.Search).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})

	http.ListenAndServe(":" + cnf.Port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
