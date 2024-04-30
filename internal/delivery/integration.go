package delivery

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) integrationGetCarByID(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNum")
	if regNumber == "" {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	car, err := h.services.CarsService.IntegrationGetCarByRegNumber(regNumber)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, car)
}
