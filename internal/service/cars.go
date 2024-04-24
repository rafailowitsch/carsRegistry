package service

import (
	"carsRegistry/internal/domain"
	"carsRegistry/internal/integration"
	"carsRegistry/internal/repository"
	"fmt"
	"github.com/google/uuid"
)

type CarsService struct {
	CarsRepo        repository.Cars
	OwnersRepo      *repository.Owners
	CarsIntegration integration.CarsInfoClient
}

func NewCarsService(carsRepo repository.Cars, ownersRepo *repository.Owners, carsIntegration integration.CarsInfoClient) *CarsService {
	return &CarsService{CarsRepo: carsRepo, OwnersRepo: ownersRepo, CarsIntegration: carsIntegration}
}

func (s *CarsService) CreateCar(carInput domain.CarsInput) error {
	if c, _ := s.CarsIntegration.FetchCarInfo(carInput.RegNumber); c.RegNum == carInput.RegNumber {
		return fmt.Errorf("car with registration number %s already exists", c.RegNum)
	}

	car := &domain.Cars{
		RegNumber: carInput.RegNumber,
		Mark:      carInput.Mark,
		Model:     carInput.Model,
		Year:      carInput.Year,
	}

	if carInput.OwnerID != uuid.Nil {
		car.OwnerID = &carInput.OwnerID
	} else {
		car.OwnerID = nil
	}

	return s.CarsRepo.CreateCar(car)
}

func (s *CarsService) DeleteCar(regNumber string) error {
	return s.CarsRepo.DeleteCar(regNumber)
}

func (s *CarsService) GetCarByRegNumber(regNumber string) (*domain.Cars, error) {
	return s.CarsRepo.GetCarByRegNumber(regNumber)
}

func (s *CarsService) GetCars(carFilter domain.CarFilter, page int, pageSize int) ([]domain.Cars, error) {
	offset := (page - 1) * pageSize
	return s.CarsRepo.GetCars(carFilter, offset, pageSize)
}

func (s *CarsService) AddNewCars(regNumbers []string) error {
	for _, regNumber := range regNumbers {
		car := &domain.Cars{RegNumber: regNumber}
		if err := s.CarsRepo.CreateCar(car); err != nil {
			return err
		}
	}
	return nil
}

func (s *CarsService) UpdateCar(carInput domain.CarsInput) error {
	car := &domain.Cars{
		RegNumber: carInput.RegNumber,
		Mark:      carInput.Mark,
		Model:     carInput.Model,
		Year:      carInput.Year,
	}

	if carInput.OwnerID != uuid.Nil {
		owner, err := (*s.OwnersRepo).GetOwnerByID(carInput.OwnerID)
		if err != nil {
			return err
		}
		car.Owner = owner
		car.OwnerID = &carInput.OwnerID
	}

	return s.CarsRepo.UpdateCar(car)
}

func (s *CarsService) IntegrationGetCarByRegNumber(regNumber string) (*domain.CarsInfo, error) {
	return s.CarsIntegration.FetchCarInfo(regNumber)
}
