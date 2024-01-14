package main

import (
	"encoding/json"
	"net/http"

	"github.com/fxivan/microservicio/transactions/pkg/models"
)

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	var m models.Transaction
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	insertResult , err := app.transactions.InsertTransaction(&m)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(insertResult)

	app.infoLog.Println("Insert transaction")
}