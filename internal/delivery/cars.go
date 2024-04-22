package delivery

import (
	"carRegistry/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) createCar(w http.ResponseWriter, r *http.Request) {
	var input domain.CarsInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return

	}

	err = h.services.CarsService.CreateCar(input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, "Car created")
}

func (h *Handler) getCars(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	var filter domain.CarFilter
	filter.RegNumber = r.URL.Query().Get("regNumber")
	filter.Mark = r.URL.Query().Get("mark")
	filter.Model = r.URL.Query().Get("model")
	filter.Year = r.URL.Query().Get("year")

	cars, err := h.services.CarsService.GetCars(filter, page, pageSize)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, cars)
}
