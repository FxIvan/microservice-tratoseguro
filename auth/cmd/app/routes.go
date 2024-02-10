package main

import (
	"net/http"

	"github.com/fxivan/microservicio/auth/pkg/middleware"
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
	//r.Handle("/auth/info", middleware.AuthMiddleware(http.HandlerFunc(app.personalInformation))).Methods("POST")
	r.Handle("/auth/info/files", middleware.AuthMiddleware(http.HandlerFunc(app.uploadFiles))).Methods("POST")

	return r
}
