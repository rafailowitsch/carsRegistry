package repository

import (
	"automobileRegistry_/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOwner(db *gorm.DB, person *domain.Owners) error {
	return db.Create(person).Error
}

func GetOwnerByID(db *gorm.DB, id uuid.UUID) (domain.Owners, error) {
	var person domain.Owners
	result := db.First(&person, "id = ?", id)
	return person, result.Error
}

func UpdateOwner(db *gorm.DB, person *domain.Owners) error {
	return db.Save(person).Error
}

func DeleteOwner(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(&domain.Owners{}, id).Error
}
