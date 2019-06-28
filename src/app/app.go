package app

import (
	"database/sql"
	"os"

	"github.com/gorilla/mux"
	"github.com/tonyooi/go-base-uservice/src/di"
)

// App - struct for app
type App struct {
	AppRouter *mux.Router
	AppDb     *sql.DB
	AppDI     di.DI
	AppLog    *os.File
}
