package delivery

import (
	"carsRegistry/internal/domain"
	l "carsRegistry/pkg/logg"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	lr "github.com/sirupsen/logrus"
	"net/http"

	"strconv"
)

func (h *Handler) createCar(w http.ResponseWriter, r *http.Request) {
	var input domain.CarsInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		l.LogError("delivery:createCar", "Failed to decode request body", err, lr.Fields{})
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	l.LogInfo("delivery:createCar", "Creating car", lr.Fields{"regNumber": input.RegNumber})
	err = h.services.CarsService.CreateCar(input)
	if err != nil {
		l.LogError("delivery:createCar", "Failed to create car", err, lr.Fields{})
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	l.LogInfo("delivery:createCar", "Car created", lr.Fields{"regNumber": input.RegNumber})

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

	l.LogInfo("delivery:getCars", "Fetching cars", lr.Fields{"page": page, "pageSize": pageSize})
	cars, err := h.services.CarsService.GetCars(filter, page, pageSize)
	if err != nil {
		l.LogError("delivery:getCars", "Failed to get cars", err, lr.Fields{})
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	l.LogInfo("delivery:getCars", "Cars fetched", lr.Fields{})

	respondWithJSON(w, http.StatusOK, cars)
}

func (h *Handler) addCars(w http.ResponseWriter, r *http.Request) {
	var request struct {
		RegNums []string `json:"regNums"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		l.LogError("delivery:addCars", "Failed to decode request body", err, lr.Fields{})
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	l.LogInfo("delivery:addCars", "Adding cars", lr.Fields{"regNums": request.RegNums})
	err = h.services.CarsService.AddNewCars(request.RegNums)
	if err != nil {
		l.LogError("delivery:addCars", "Failed to add cars", err, lr.Fields{})
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	l.LogInfo("delivery:addCars", "Cars added", lr.Fields{"regNums": request.RegNums})

	respondWithJSON(w, http.StatusCreated, "Car created")
}

func (h *Handler) updateCar(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNumber")
	if regNumber == "" {
		l.LogError("delivery:updateCar", "Invalid request payload", nil, lr.Fields{})
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var input domain.CarsInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		l.LogError("delivery:updateCar", "Failed to decode request body", err, lr.Fields{})
		respondWithJSON(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	l.LogInfo("delivery:updateCar", "Updating car", lr.Fields{"regNumber": regNumber})
	err = h.services.CarsService.UpdateCar(input)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	l.LogInfo("delivery:updateCar", "Car updated", lr.Fields{"regNumber": regNumber})

	respondWithJSON(w, http.StatusOK, "Car updated")
}

func (h *Handler) getCarByID(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNumber")
	if regNumber == "" {
		l.LogError("delivery:getCarByID", "Invalid request payload", nil, lr.Fields{})
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	l.LogInfo("delivery:getCarByID", "Fetching car", lr.Fields{"regNumber": regNumber})
	car, err := h.services.CarsService.GetCarByRegNumber(regNumber)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	l.LogInfo("delivery:getCarByID", "Car fetched", lr.Fields{"regNumber": regNumber})

	respondWithJSON(w, http.StatusOK, car)
}

func (h *Handler) deleteCar(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNumber")
	if regNumber == "" {
		l.LogError("delivery:deleteCar", "Invalid request payload", nil, lr.Fields{})
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	l.LogInfo("delivery:deleteCar", "Deleting car", lr.Fields{"regNumber": regNumber})
	err := h.services.CarsService.DeleteCar(regNumber)
	if err != nil {
		l.LogError("delivery:deleteCar", "Failed to delete car", err, lr.Fields{"regNumber": regNumber})
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	l.LogInfo("delivery:deleteCar", "Car deleted", lr.Fields{"regNumber": regNumber})

	respondWithJSON(w, http.StatusOK, "Car deleted")
}
