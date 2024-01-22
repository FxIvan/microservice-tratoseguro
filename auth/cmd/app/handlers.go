package main

import (
	"encoding/json"
	"net/http"

	"github.com/fxivan/microservicio/auth/pkg/models"
)
func (app *application) insert(w http.ResponseWriter, r *http.Request){
	var m models.UserSingIn

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil{
		app.errorLog.Println(err)
	}

	insertResult, err := app.users.InsertRegisterUser(&m)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(insertResult)

}