package repository

import (
	"automobileRegistry_/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCar(db *gorm.DB, car *domain.Cars) error {
	return db.Create(car).Error
}

func GetCarByID(db *gorm.DB, id uuid.UUID) (domain.Cars, error) {
	var car domain.Cars
	result := db.Preload("Owner").First(&car, "id = ?", id)
	return car, result.Error
}

func UpdateCar(db *gorm.DB, car *domain.Cars) error {
	return db.Save(car).Error
}

func DeleteCar(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(&domain.Cars{}, id).Error
}
