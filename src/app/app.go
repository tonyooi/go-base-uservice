package app

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/tonyooi/go-base-uservice/src/di"
)

// App - struct for app
type App struct {
	AppRouter *mux.Router
	AppDb     *sql.DB
	AppDI     di.DI
}
