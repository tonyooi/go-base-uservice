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
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	var app = new(app.App)
	var wait time.Duration
	var server *http.Server
	var file *os.File
	var err error

	app.AppRouter = mux.NewRouter()

	if file, err = os.OpenFile("logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
		log.Println(err.Error())
		log.Fatal("Cannot open log file. Server shutting down.")
	}
	file.Close()

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/log.txt",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

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
