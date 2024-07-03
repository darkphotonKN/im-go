package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/darkphotonKN/go-im/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
