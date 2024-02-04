package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()

	//Rutas publica
	//Registro
	r.HandleFunc("/auth/signup", app.signup).Methods("POST")
	//Login
	r.HandleFunc("/auth/signin", app.signin).Methods("POST")

	//Routes with middleware PROTECTED
	//Example. Importar middleware que esta dentro de pkg
	//r.Handle("/api/signup", middleware.AuthMiddleware(http.HandlerFunc(app.signup))).Methods("POST")

	return r
}
