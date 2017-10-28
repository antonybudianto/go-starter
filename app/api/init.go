package api

import (
	"github.com/antonybudianto/go-starter/app/api/handler"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Register api routes here
func RegisterRoutes(r *mux.Router) {
	handler.RegisterPersonRoutes(r)
}
