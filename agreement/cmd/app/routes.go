package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	//r.Handle("/load/img", middleware.AuthMiddleware(http.HandlerFunc(app.uploadImg))).Methods("POST")
	//r.Handle("/load/document", middleware.AuthMiddleware(http.HandlerFunc(app.uploadFile))).Methods("POST")
	r.HandleFunc("/search/user", app.searchCTPY).Methods("POST")
	return r
}
