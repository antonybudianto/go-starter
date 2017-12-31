package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/antonybudianto/go-starter/routes/user"
)

// App level struct containing its dependencies
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize App dependencies
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	u := user.UserHandler{a.Router, a.DB}
	u.InitializeRoutes()
}
