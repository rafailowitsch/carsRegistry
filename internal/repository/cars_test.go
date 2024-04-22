package repository

import (
	"automobileRegistry_/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestCreateAndUpdateCar(t *testing.T) {
	owner := domain.Owners{
		ID:      uuid.New(),
		Name:    "Bob",
		Surname: "Jones",
	}
	_ = CreateOwner(db, &owner)

	car := domain.Cars{
		BaseModel: domain.BaseModel{
			ID: uuid.New(),
		},
		RegNumber: "X999XX99",
		Mark:      "Toyota",
		Model:     "Camry",
		Year:      "2020",
		OwnerID:   owner.ID,
	}
	err := CreateCar(db, &car)
	if err != nil {
		t.Fatalf("Failed to create car: %v", err)
	}

	car.Mark = "Honda"
	err = UpdateCar(db, &car)
	if err != nil {
		t.Errorf("Failed to update car: %v", err)
	}

	retrievedCar, err := GetCarByID(db, car.ID)
	if err != nil {
		t.Errorf("Failed to retrieve car: %v", err)
	}

	if retrievedCar.Mark != "Honda" {
		t.Errorf("Retrieved car mark mismatch: expected Honda, got %s", retrievedCar.Mark)
	}

	_ = DeleteCar(db, car.ID)     // Очистка после теста
	_ = DeleteOwner(db, owner.ID) // Удаление владельца
}
