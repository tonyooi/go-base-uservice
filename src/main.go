package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/tonyooi/go-base-uservice/src/app"
)

func main() {

	var app = new(app.App)
	var wait time.Duration
	var server *http.Server

	app.AppRouter = mux.NewRouter()

	server = &http.Server{
		Addr:           ":9000",
		Handler:        app.AppRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// add dependecny injection to app.AppDI
	app.InitDependencyInjection()

	// add global middlewares
	app.InitGlobalMiddlewares()

	// add routes to app
	app.InitRoutes()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Server gracefully shutting down")
	os.Exit(0)
}
