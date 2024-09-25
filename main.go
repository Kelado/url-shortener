package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    Addr,
		Handler: r,
	}

	go activateGracefullShutdown(server)

	log.Println("Listerning on", Addr)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

	log.Println("Grecafull Shutdown")
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

func activateGracefullShutdown(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
}
