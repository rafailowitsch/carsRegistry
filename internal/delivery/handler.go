package delivery

import (
	"carRegistry/internal/service"
	"encoding/json"
	"net/http"
)
import "github.com/go-chi/chi/v5"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/owners", func(r chi.Router) {
		r.Post("/", h.createOwner)
		//r.Get("/{id}", h.getOwnerByID)
		//r.Put("/{id}", h.updateOwner)
		//r.Delete("/{id}", h.deleteOwner)
	})

	r.Route("/cars", func(r chi.Router) {
		r.Post("/", h.createCar)
		//r.Get("/{id}", h.getCarByID)
		//r.Put("/{id}", h.updateCar)
		//r.Delete("/{id}", h.deleteCar)
		r.Get("/getCars", h.getCars)
	})

	return r
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
