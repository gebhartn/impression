package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gebhartn/impression/files"
	"github.com/gebhartn/impression/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var bindAddress = getEnv("BIND_ADDRESS", ":9090")
var logLevel = getEnv("LOG_LEVEL", "debug")

// todo: we'll use this as a temp store on the filesystem until an s3 implementation is in
var basePath = getEnv("BASE_PATH", "./imagestore")

func main() {

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "images-service",
			Level: hclog.LevelFromString(logLevel),
			Color: hclog.AutoColor,
		},
	)

	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	store, err := files.NewLocal(basePath, 1024*1000*5)
	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	fh := handlers.NewFiles(store, l)
	mw := handlers.GzipHandler{}

	sm := mux.NewRouter()
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	ph := sm.Methods(http.MethodPut).Subrouter()
	// todo: route upload files
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+}", fh.UploadREST)

	gh := sm.Methods(http.MethodGet).Subrouter()
	// todo: route get files
	gh.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+}", http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))
	gh.Use(mw.GzipMiddleware)

	s := http.Server{
		Addr:         bindAddress,
		Handler:      ch(sm),
		ErrorLog:     sl,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Info("Starting the server", "bind_address", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := <-c

	l.Info("Shutting down server with", "signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	s.Shutdown(ctx)
}
