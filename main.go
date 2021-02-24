package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// todo: initialize the store w/ max file size
	// todo: create handlers/attach [new files, gzip]

	sm := mux.NewRouter()
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// todo: route upload files
	// todo: route get files

	// todo: compression

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
