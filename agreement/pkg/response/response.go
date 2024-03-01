package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func HttpResponseError(w http.ResponseWriter, responeStruct *Response) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responeStruct); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
