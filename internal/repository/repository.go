package repository

import (
	"carsRegistry/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Owners interface {
	CreateOwner(owner *domain.Owners) error
	GetOwnerByID(id uuid.UUID) (*domain.Owners, error)
	UpdateOwner(owner *domain.Owners) error
	DeleteOwner(id uuid.UUID) error
}

type Cars interface {
	CreateCar(car *domain.Cars) error
	GetCarByRegNumber(regNumber string) (*domain.Cars, error)
	UpdateCar(car *domain.Cars) error
	DeleteCar(regNumber string) error
	GetCars(filter domain.CarFilter, offset int, limit int) ([]domain.Cars, error)
}

type Repository struct {
	OwnersRepo Owners
	CarsRepo   Cars
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		OwnersRepo: NewOwnersRepo(db),
		CarsRepo:   NewCarsRepo(db),
	}
}
