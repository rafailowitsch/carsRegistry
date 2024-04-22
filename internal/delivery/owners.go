package delivery

import (
	"carRegistry/internal/domain"
	"encoding/json"
	"net/http"
)

func (h *Handler) createOwner(w http.ResponseWriter, r *http.Request) {
	var input domain.OwnersInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.services.OwnersService.CreateOwner(input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, "Owner created")
}
