package handler

import (
	"encoding/json"
	"net/http"

	controller "github.com/Kelado/url-shortener/controllers"
	"github.com/Kelado/url-shortener/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	controller *controller.Controller
}

func NewHandler(controller *controller.Controller) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (h *Handler) HandlePostURL(w http.ResponseWriter, r *http.Request) {
	var link models.LinkRequest
	json.NewDecoder(r.Body).Decode(&link)

	if err := link.ValidateSchema(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	shortenedURL, err := h.controller.CreateLink(link)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	resp := models.LinkResponse{
		ShortURL: shortenedURL,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}

func (h *Handler) HandleGetURL(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	originalURL, err := h.controller.GetLink(code)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, string(originalURL), http.StatusSeeOther)
}
