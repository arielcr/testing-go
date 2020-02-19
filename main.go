package main

import (
	"github.com/arielcr/testing-go/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	//total := math.Add(1, 2)
	//fmt.Printf("The result is: %v\n", total)

	router := mux.NewRouter()
	router.HandleFunc("/person/create", api.CreatePersonEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))

}