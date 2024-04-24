package repository

import (
	"carsRegistry/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type CarsRepo struct {
	db *gorm.DB
}

func NewCarsRepo(db *gorm.DB) *CarsRepo {
	return &CarsRepo{db: db}
}

func (c *CarsRepo) CreateCar(car *domain.Cars) error {
	return c.db.Create(car).Error
}

func (c *CarsRepo) GetCarByRegNumber(regNumber string) (*domain.Cars, error) {
	var car domain.Cars
	result := c.db.Preload("Owner").First(&car, "reg_number = ?", regNumber)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println(car)
	return &car, nil
}

func (c *CarsRepo) UpdateCar(car *domain.Cars) error {
	existingCar, err := c.GetCarByRegNumber(car.RegNumber)
	if err != nil {
		return err
	}

	existingCar.Mark = car.Mark
	existingCar.Model = car.Model
	existingCar.Year = car.Year
	if *existingCar.OwnerID != uuid.Nil {
		existingCar.OwnerID = car.OwnerID
		existingCar.Owner = car.Owner
	}
	log.Println(existingCar, *existingCar.Owner)
	return c.db.Save(existingCar).Error
}

func (c *CarsRepo) DeleteCar(regNumber string) error {
	return c.db.Where("reg_number = ?", regNumber).Delete(&domain.Cars{}).Error
}

func (c *CarsRepo) GetCars(filter domain.CarFilter, offset int, limit int) ([]domain.Cars, error) {
	var cars []domain.Cars
	query := c.db.Preload("Owner").Model(&domain.Cars{})
	query = filterCars(query, filter)

	query = query.Offset(offset).Limit(limit)
	if err := query.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func filterCars(query *gorm.DB, filter domain.CarFilter) *gorm.DB {
	if filter.RegNumber != "" {
		query = query.Where("reg_number = ?", filter.RegNumber)
	}
	if filter.Mark != "" {
		query = query.Where("mark = ?", filter.Mark)
	}
	if filter.Model != "" {
		query = query.Where("model = ?", filter.Model)
	}
	if filter.Year != "" {
		query = query.Where("year = ?", filter.Year)
	}
	if filter.OwnerID != uuid.Nil {
		query = query.Where("owner_id = ?", filter.OwnerID)
	}
	return query
}
