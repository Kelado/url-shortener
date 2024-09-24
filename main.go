package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	Addr = ":8000"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", HandlePing)

	r.Post("/shorten", HandlePostURL)
	r.Get("/:shortUrl", HandleGetURL)

	http.ListenAndServe(Addr, r)
}
