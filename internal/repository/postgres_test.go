package repository

import (
	"carRegistry/internal/domain"
	"log"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	dsn := "host=localhost user=postgres password=qwerty dbname=carRegistry port=5444"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	_ = db.AutoMigrate(&domain.Owners{}, &domain.Cars{})

	code := m.Run()
	os.Exit(code)
}
