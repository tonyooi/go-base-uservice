package app

import (
	"github.com/tonyooi/go-base-uservice/src/globalmiddlewares"
)

// InitGlobalMiddlewares - define all global middlewares here
func (app *App) InitGlobalMiddlewares() {
	app.AppRouter.Use(globalmiddlewares.Authenticate)
}
