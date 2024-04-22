package service

import (
	"carsRegistry/internal/domain"
	"carsRegistry/internal/repository"
)

type CarsService struct {
	CarsRepo repository.Cars
}

func NewCarsService(carsRepo repository.Cars) *CarsService {
	return &CarsService{CarsRepo: carsRepo}
}

func (s *CarsService) CreateCar(carInput domain.CarsInput) error {
	car := &domain.Cars{
		RegNumber: carInput.RegNumber,
		Mark:      carInput.Mark,
		Model:     carInput.Model,
		Year:      carInput.Year,
		OwnerID:   carInput.OwnerID,
	}

	return s.CarsRepo.CreateCar(car)
}

func (s *CarsService) GetCars(carFilter domain.CarFilter, page int, pageSize int) ([]domain.Cars, error) {
	offset := (page - 1) * pageSize
	return s.CarsRepo.GetCars(carFilter, offset, pageSize)
}
