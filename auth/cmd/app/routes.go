package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/api/singin", app.insert).Methods("POST")
	return r
}