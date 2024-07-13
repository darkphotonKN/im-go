package main

import (
	"net/http"

	"github.com/darkphotonKN/go-im/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/ws", handlers.WsEndpoint)

	return mux
}
