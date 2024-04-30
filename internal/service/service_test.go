package service

import (
	"carsRegistry/internal/domain"
	"carsRegistry/internal/integration"
	"carsRegistry/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCarsService_CreateCar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCarsRepo := mock.NewMockCars(ctrl)
	mockOwnersRepo := mock.NewMockOwners(ctrl)

	carsService := NewCarsService(mockCarsRepo, mockOwnersRepo, integration.CarsInfoClient{})

	carInput := domain.CarsInput{
		RegNumber: "X999XX99",
		Mark:      "Toyota",
		Model:     "Camry",
		Year:      "2020",
	}

	mockCarsRepo.EXPECT().CreateCar(gomock.Any()).Return(nil)

	err := carsService.CreateCar(carInput)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
