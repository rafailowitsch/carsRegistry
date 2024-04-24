package delivery

import (
	"carsRegistry/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
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
	log.Println("Request createCar:", input)

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

func (h *Handler) addCars(w http.ResponseWriter, r *http.Request) {
	var request struct {
		RegNums []string `json:"regNums"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.services.CarsService.AddNewCars(request.RegNums)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, "Car created")
}

func (h *Handler) updateCar(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNumber")
	if regNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Println(regNumber)

	var input domain.CarsInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Println(input.RegNumber)
	err = h.services.CarsService.UpdateCar(input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Car updated")
}

func (h *Handler) getCarByID(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNumber")
	if regNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	car, err := h.services.CarsService.GetCarByRegNumber(regNumber)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, car)
}

func (h *Handler) deleteCar(w http.ResponseWriter, r *http.Request) {
	regNumber := chi.URLParam(r, "regNumber")
	if regNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := h.services.CarsService.DeleteCar(regNumber)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Car deleted")
}
