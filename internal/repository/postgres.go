package repository

import (
	"automobileRegistry_/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.Cars{}, &domain.Owners{})
	if err != nil {
		return nil, err
	}
	return db, nil
}