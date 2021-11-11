package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hpcsc/argocd-playground-app/handlers"
	"github.com/hpcsc/argocd-playground-app/middlewares"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	env := os.Getenv("ENVIRONMENT")

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar().With(zap.String("env", env))

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	r.Handle("/info", handlers.NewInfoHandler("version.json"))

	r.Use(middlewares.WithHttpContext)

	address := ":8888"
	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	sugar.Info(fmt.Sprintf("starting server at %s", address))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			sugar.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	<-c

	wait := 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	sugar.Info("shutting down")
	os.Exit(0)
}
