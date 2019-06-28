package app

import (
	"github.com/tonyooi/go-base-uservice/src/middlewares"
)

// InitRoutes - all routes handler functions here
func (app *App) InitRoutes() {
	// add all routes here

	/* "/" */
	app.AppRouter.HandleFunc("/", middlewares.Logging(app.GetHome)).Methods("GET")
	// app.AppRouter.HandleFunc("/", home.GetHome).Methods("GET")
}
