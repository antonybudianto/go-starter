package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/antonybudianto/go-starter/app/api/models"
)

var people = []models.Person {
	models.Person{
		ID: "1",
		Firstname: "Antony",
		Lastname: "Budianto",
		Address: &models.Address{City: "City X", State: "State X"},
	},
}

// GetPeople - Display all from the people var
func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson - Display a single data
func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
			if item.ID == params["id"] {
					json.NewEncoder(w).Encode(item)
					return
			}
	}
	json.NewEncoder(w).Encode(&models.Person{})
}

// CreatePerson - create a new item
func createPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson - Delete an item
func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
			if item.ID == params["id"] {
					people = append(people[:index], people[index+1:]...)
					break
			}
			json.NewEncoder(w).Encode(people)
	}
}

// RegisterPersonRoutes - Register api routes here
func RegisterPersonRoutes(r *mux.Router) {
	r.HandleFunc("/people", getPeople).Methods("GET")
	r.HandleFunc("/people/{id}", getPerson).Methods("GET")
	r.HandleFunc("/people/{id}", createPerson).Methods("POST")
	r.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")
}
