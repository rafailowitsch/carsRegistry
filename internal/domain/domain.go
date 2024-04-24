package domain

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Cars struct {
	BaseModel
	RegNumber string     `gorm:"primaryKey;not null"`
	Mark      string     `gorm:"not null"`
	Model     string     `gorm:"not null"`
	Year      string     `gorm:"not null"`
	OwnerID   *uuid.UUID `gorm:"type:uuid"`
	Owner     *Owners    `gorm:"foreignKey:OwnerID"`
}

type Owners struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name       string    `gorm:"not null"`
	Surname    string    `gorm:"not null"`
	Patronymic string
	Cars       []Cars `gorm:"foreignKey:OwnerID"`
}

type CarsInput struct {
	RegNumber string    `json:"regNumber"`
	Mark      string    `json:"mark"`
	Model     string    `json:"model"`
	Year      string    `json:"year"`
	OwnerID   uuid.UUID `json:"ownerID"`
}

type OwnersInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type CarFilter struct {
	RegNumber string
	Mark      string
	Model     string
	Year      string
	OwnerID   uuid.UUID
}

type Pagination struct {
	Page     int
	PageSize int
	Offset   int
}

type CarsInfo struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
	} `json:"owner"`
}
