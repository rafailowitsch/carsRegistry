package repository

import (
	"carRegistry/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OwnersRepo struct {
	db *gorm.DB
}

func NewOwnersRepo(db *gorm.DB) *OwnersRepo {
	return &OwnersRepo{db: db}
}

func (o *OwnersRepo) CreateOwner(owner *domain.Owners) error {
	return o.db.Create(owner).Error
}

func (o *OwnersRepo) GetOwnerByID(id uuid.UUID) (*domain.Owners, error) {
	var owner domain.Owners
	result := o.db.First(&owner, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &owner, nil
}

func (o *OwnersRepo) UpdateOwner(owner *domain.Owners) error {
	return o.db.Save(owner).Error
}

func (o *OwnersRepo) DeleteOwner(id uuid.UUID) error {
	return o.db.Delete(&domain.Owners{}, id).Error
}
