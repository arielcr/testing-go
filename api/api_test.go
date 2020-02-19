package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/person/create", CreatePersonEndpoint).Methods("POST")

	return router

}

func TestPersonEndpoint(t *testing.T) {

	person := &Person{
		FirstName: "Ariel",
		LastName:  "Orozco",
	}

	jsonPerson, _ := json.Marshal(person)

	request, _ := http.NewRequest("POST", "/person/create", bytes.NewBuffer(jsonPerson))

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

}

func TestPersonEndpointFailed(t *testing.T) {

	request, _ := http.NewRequest("POST", "/person/create", bytes.NewBuffer([]byte("Hello")))

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)

}