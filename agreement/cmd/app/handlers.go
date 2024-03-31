package main

import (
	"encoding/json"
	"fmt"
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

func (app *application) createContractPRNE(w http.ResponseWriter, r *http.Request) {
	var m models.ContractDefinitionModel

	email, ok := r.Context().Value("email").(string)

	if !ok {
		responseErr := &response.Response{
			Status:  false,
			Message: "Error al obtener informacion del usuario",
			Code:    401,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseErr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		responsErr := &response.Response{
			Status:  false,
			Message: fmt.Sprintf("Error al recibir la informacion %s", err),
			Code:    400,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responsErr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	responseMessage, status := app.agreement.SaveAgreement(&m, email)

	if status == false {
		responseErr := &response.Response{
			Status:  false,
			Message: responseMessage,
			Code:    400,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseErr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
