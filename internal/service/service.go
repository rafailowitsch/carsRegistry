package service

import (
	"carsRegistry/internal/domain"
	"carsRegistry/internal/integration"
	"carsRegistry/internal/repository"
)

type Owners interface {
	CreateOwner(ownerInput domain.OwnersInput) error
	GetOwnerByID(id string) (*domain.Owners, error)
	//UpdateOwner(id string, name, surname, patronymic string) error
	//DeleteOwner(id string) error
}

type Cars interface {
	CreateCar(carInput domain.CarsInput) error
	GetCarByRegNumber(regNumber string) (*domain.Cars, error)
	UpdateCar(carInput domain.CarsInput) error
	DeleteCar(regNumber string) error
	AddNewCars(regNumbers []string) error
	GetCars(carFilter domain.CarFilter, page int, pageSize int) ([]domain.Cars, error)
	IntegrationGetCarByRegNumber(regNumber string) (*domain.CarsInfo, error)
}

type Services struct {
	OwnersService Owners
	CarsService   Cars
}

func NewServices(repo *repository.Repository, integration *integration.Integrations) *Services {
	ownersService := NewOwnersService(repo.OwnersRepo)
	return &Services{
		OwnersService: ownersService,
		CarsService:   NewCarsService(repo.CarsRepo, &ownersService.OwnersRepo, *integration.CarsInfoClient),
	}
}
