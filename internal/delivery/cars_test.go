package delivery

import (
	"bytes"
	"carsRegistry/internal/domain"
	"carsRegistry/internal/service"
	mock_service "carsRegistry/internal/service/mocks"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateCar(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCars, input domain.CarsInput)

	tests := []struct {
		name         string
		body         string
		input        domain.CarsInput
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name: "ok",
			body: `{"regNumber": "X999XX99", "mark": "Toyota", "model": "Camry", "year": "2020"}`,
			input: domain.CarsInput{
				RegNumber: "X999XX99",
				Mark:      "Toyota",
				Model:     "Camry",
				Year:      "2020",
			},
			mockBehavior: func(r *mock_service.MockCars, input domain.CarsInput) {
				r.EXPECT().CreateCar(input).Return(nil)
			},
			statusCode:   201,
			responseBody: "\"Car created\"",
		},
		{
			name: "invalid input",
			body: `{wrong}`,
			mockBehavior: func(r *mock_service.MockCars, input domain.CarsInput) {
			},
			statusCode:   400,
			responseBody: "\"Invalid request payload\"",
		},
		{
			name: "service error",
			body: `{"regNumber": "X999XX99", "mark": "Toyota", "model": "Camry", "year": "2020"}`,
			input: domain.CarsInput{
				RegNumber: "X999XX99",
				Mark:      "Toyota",
				Model:     "Camry",
				Year:      "2020",
			},
			mockBehavior: func(r *mock_service.MockCars, input domain.CarsInput) {
				r.EXPECT().CreateCar(input).Return(errors.New("failed to create car"))
			},
			statusCode:   500,
			responseBody: "\"failed to create car\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockCars(c)
			tt.mockBehavior(s, tt.input)

			services := &service.Services{CarsService: s}
			handler := Handler{services: services}

			// Init Endpoint
			r := chi.NewRouter()
			r.Post("/cars", handler.createCar)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/cars",
				bytes.NewBufferString(tt.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.statusCode, w.Code)
			assert.Equal(t, tt.responseBody, w.Body.String())
		})
	}
}

func TestHandler_UpdateCar(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCars, input domain.CarsInput)

	tests := []struct {
		name         string
		regNumber    string
		body         string
		input        domain.CarsInput
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name:      "ok",
			regNumber: "X999XX99",
			body:      `{"mark": "Toyota", "model": "Camry", "year": "2020"}`,
			input: domain.CarsInput{
				Mark:  "Toyota",
				Model: "Camry",
				Year:  "2020",
			},
			mockBehavior: func(r *mock_service.MockCars, input domain.CarsInput) {
				r.EXPECT().UpdateCar(input).Return(nil)
			},
			statusCode:   200,
			responseBody: "\"Car updated\"",
		},
		//{
		//	name: "service error",
		//	body: `{"regNumber": "X999XX99", "mark": "Toyota", "model": "Camry", "year": "2020"}`,
		//	input: domain.CarsInput{
		//		RegNumber: "X999XX99",
		//		Mark:      "Toyota",
		//		Model:     "Camry",
		//		Year:      "2020",
		//	},
		//	mockBehavior: func(r *mock_service.MockCars, input domain.CarsInput) {
		//		r.EXPECT().UpdateCar(input).Return(errors.New("failed to update car"))
		//	},
		//	statusCode:   500,
		//	responseBody: "\"failed to update car\"",
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockCars(c)
			tt.mockBehavior(s, tt.input)

			services := &service.Services{CarsService: s}
			handler := Handler{services: services}

			r := chi.NewRouter()
			r.Put("/cars/"+tt.regNumber, handler.updateCar)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/cars/"+tt.regNumber,
				bytes.NewBufferString(tt.body))

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)
			assert.Equal(t, tt.responseBody, w.Body.String())
		})
	}
}
