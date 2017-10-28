package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "github.com/antonybudianto/go-starter/app/api/handler"
)

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/people", handler.GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", handler.GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", handler.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", handler.DeletePerson).Methods("DELETE")
	log.Printf("API served at localhost:%d", 8000)
    log.Fatal(http.ListenAndServe(":8000", router))
}
