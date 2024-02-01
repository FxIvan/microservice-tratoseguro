package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	//Registro
	r.HandleFunc("/api/signup", app.signup).Methods("POST")
	//Login
	r.HandleFunc("/api/signin", app.signin).Methods("POST")
	return r
}
