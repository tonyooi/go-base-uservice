package app

import (
	"github.com/tonyooi/go-base-uservice/src/di"
)

// InitDependencyInjection - add depency in jections here
func (app *App) InitDependencyInjection() {
	app.AppDI = di.InitDI()
}
