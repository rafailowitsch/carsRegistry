package delivery

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) integrationGetCarByID(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNum")
	if regNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	car, err := h.services.CarsService.IntegrationGetCarByRegNumber(regNumber)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, car)
}
