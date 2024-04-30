package delivery

import (
	"carsRegistry/internal/domain"
	l "carsRegistry/pkg/logg"
	"encoding/json"
	lr "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) createOwner(w http.ResponseWriter, r *http.Request) {
	var input domain.OwnersInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		l.LogError("delivery:createOwner", "Failed to decode request body", err, lr.Fields{})
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	l.LogInfo("delivery:createOwner", "Creating owner", lr.Fields{"name": input.Name})
	err = h.services.OwnersService.CreateOwner(input)
	if err != nil {
		l.LogError("delivery:createOwner", "Failed to create owner", err, lr.Fields{})
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	l.LogInfo("delivery:createOwner", "Owner created", lr.Fields{"name": input.Name})

	respondWithJSON(w, http.StatusCreated, "Owner created")
}
