package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fxivan/microservicio/agreement/pkg/models"
)

func (app *application) searchCTPY(w http.ResponseWriter, r *http.Request) {

	var m models.SearchUser
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		panic(err)
		return
	}

	allUsers := app.
	return
}
