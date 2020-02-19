package api

import (
	"encoding/json"
	"net/http"
)

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	var p Person

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(Response{
		Status: true,
		Code: 200,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(response)

}