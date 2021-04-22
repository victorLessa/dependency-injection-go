package app

import (
	"log"
	"net/http"
	database "webserver/app/Config"

	migrations "webserver/app/Database/migrations"
	routes "webserver/app/Routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	
	migrations.RunMigrations()
	a.DB = database.Connect()
	a.Router = mux.NewRouter()
	
	routes.Book(a.Router, a.DB)
	routes.Book(a.Router, a.DB)

}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}