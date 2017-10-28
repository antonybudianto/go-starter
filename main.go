package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "github.com/antonybudianto/go-starter/app/api"
)

func main() {
    router := mux.NewRouter()
    api.RegisterRoutes(router)

    log.Printf("API served at localhost:%d", 8000)
    log.Fatal(http.ListenAndServe(":8000", router))
}
