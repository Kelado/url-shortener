package main

import (
	"net/http"

	controller "github.com/Kelado/url-shortener/controllers"
	handler "github.com/Kelado/url-shortener/handlers"
	"github.com/Kelado/url-shortener/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	Addr     = ":8000"
	CodeSize = 6
	Hostname = "http://localhost:8000/"
)

func main() {
	linkRepo := repositories.NewMockDB()
	controller := controller.NewController(Hostname, CodeSize, linkRepo)
	handler := handler.NewHandler(controller)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", handler.HandlePing)

	r.Post("/shorten", handler.HandlePostURL)
	r.Get("/:shortUrl", handler.HandleGetURL)

	http.ListenAndServe(Addr, r)
}
