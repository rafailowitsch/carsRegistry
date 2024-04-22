package domain

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Cars struct {
	BaseModel
	RegNumber string    `gorm:"not null"`
	Mark      string    `gorm:"not null"`
	Model     string    `gorm:"not null"`
	Year      string    `gorm:"not null"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null"`
	Owner     Owners    `gorm:"foreignKey:OwnerID"`
}

type Owners struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Name       string    `gorm:"not null"`
	Surname    string    `gorm:"not null"`
	Patronymic string
	Cars       []Cars `gorm:"foreignKey:OwnerID"`
}
