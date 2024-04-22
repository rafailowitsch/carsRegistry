package service

import (
	"carRegistry/internal/domain"
	"carRegistry/internal/repository"
)

type Owners interface {
	CreateOwner(ownerInput domain.OwnersInput) error
	//GetOwnerByID(id string) (Owners, error)
	//UpdateOwner(id string, name, surname, patronymic string) error
	//DeleteOwner(id string) error
}

type Cars interface {
	CreateCar(carInput domain.CarsInput) error
	//GetCarByID(id string) (Cars, error)
	//UpdateCar(id, brand, model, color, ownerID string) error
	//DeleteCar(id string) error
	GetCars(carFilter domain.CarFilter, page int, pageSize int) ([]domain.Cars, error)
}

type Services struct {
	OwnersService Owners
	CarsService   Cars
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		OwnersService: NewOwnersService(repo.OwnersRepo),
		CarsService:   NewCarsService(repo.CarsRepo),
	}
}
