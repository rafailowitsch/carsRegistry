package repository

import (
	"carsRegistry/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		//Logger: logg.go.Default.LogMode(logg.go.Info),
	})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.Cars{}, &domain.Owners{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
