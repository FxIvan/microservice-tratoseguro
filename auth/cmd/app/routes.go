package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/signup", app.insert).Methods("POST")
	r.HandleFunc("/api/signin", app.signin).Methods("POST")
	return r
}
