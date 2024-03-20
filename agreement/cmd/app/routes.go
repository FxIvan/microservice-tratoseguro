package main

import (
	"net/http"

	"github.com/fxivan/microservicio/agreement/pkg/middleware"
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/search/user", app.searchCTPY).Methods("POST")
	r.Handle("/create/agreement", middleware.AuthMiddleware(http.HandlerFunc(app.createContractPRNE))).Methods("POST")
	return r
}
