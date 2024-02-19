package main

import (
	"net/http"

	"github.com/fxivan/microservicio/media/pkg/middleware"
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/load/img", middleware.AuthMiddleware(http.HandlerFunc(app.uploadImg))).Methods("POST")
	return r
}
