package main

import (
	"encoding/json"
	"net/http"

	"github.com/fxivan/microservicio/agreement/pkg/models"
	"github.com/fxivan/microservicio/agreement/pkg/response"
)

func (app *application) searchCTPY(w http.ResponseWriter, r *http.Request) {

	var m models.SearchUser
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		panic(err)
		return
	}

	responseSearch, errSearch := app.users.SearchUser(&m)

	if errSearch == false {
		responseSucc := &response.Response{
			Status:  true,
			Message: responseSearch,
			Code:    200,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseSucc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	responseSucc := &response.Response{
		Status:  true,
		Message: responseSearch,
		Code:    200,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseSucc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
