package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	controller "github.com/Kelado/url-shortener/controllers"
	handler "github.com/Kelado/url-shortener/handlers"
	"github.com/Kelado/url-shortener/repositories"
)

var (
	Addr     = ":8000"
	CodeSize = 6
	Hostname = "http://localhost:8000/"
)

func main() {
	linkRepo := repositories.NewSQLiteDB(nil)
	controller := controller.NewController(Hostname, CodeSize, linkRepo)
	handler := handler.NewHandler(controller)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", handler.HandlePing)

	r.Post("/shorten", handler.HandlePostURL)
	r.Get("/{code}", handler.HandleGetURL)

	log.Println("Listerning on", Addr)
	http.ListenAndServe(Addr, r)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hostname := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")

	Addr = ":" + port
	Hostname = hostname + Addr + "/"
}
