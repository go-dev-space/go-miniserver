package main

import (
	"net/http"
	"time"

	"github.com/go-dev-space/go-miniserver/internal/handlers"
)

type system struct {
	Config   *config
	Handlers *handlers.RouteHandlers
}

type config struct {
	Addr string
}

func (sys system) Mount() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", sys.Handlers.HealthCheck)

	return mux
}

func (sys system) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         sys.Config.Addr,
		Handler:      mux,
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return srv.ListenAndServe()
}
