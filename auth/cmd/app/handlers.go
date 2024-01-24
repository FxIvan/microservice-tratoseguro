package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fxivan/microservicio/auth/pkg/models"
)

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	var m models.UserSignup
	fmt.Println("Insertando Usuario", m)
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
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

func (app *application) signin(w http.ResponseWriter, r *http.Request) {
	var m models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.errorLog.Println(err)
	}
	userSingIn, err := app.users.FindUserEmail(m.Username)
	fmt.Println("Informacion del usuario: ", userSingIn)
}
