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
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson - Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
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
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson - Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
			if item.ID == params["id"] {
					people = append(people[:index], people[index+1:]...)
					break
			}
			json.NewEncoder(w).Encode(people)
	}
}
