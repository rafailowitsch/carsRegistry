package repository

import (
	"carsRegistry/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestCreateAndUpdateCar(t *testing.T) {
	owner := domain.Owners{
		ID:         uuid.New(),
		Name:       "Bob",
		Surname:    "Jones",
		Patronymic: "Middle",
	}

	ownersRepo := NewOwnersRepo(db)

	err := ownersRepo.CreateOwner(&owner)
	if err != nil {
		t.Fatalf("Failed to create owner: %v", err)
	}

	car := domain.Cars{
		RegNumber: "X999XX99",
		Mark:      "Toyota",
		Model:     "Camry",
		Year:      "2020",
		OwnerID:   &owner.ID,
	}

	carsRepo := NewCarsRepo(db)

	err = carsRepo.CreateCar(&car)
	if err != nil {
		t.Fatalf("Failed to create car: %v", err)
	}

	car.Mark = "Honda"
	err = carsRepo.UpdateCar(&car)
	if err != nil {
		t.Errorf("Failed to update car: %v", err)
	}

	retrievedCar, err := carsRepo.GetCarByRegNumber(car.RegNumber)
	if err != nil {
		t.Errorf("Failed to retrieve car: %v", err)
	}

	if retrievedCar.Mark != "Honda" {
		t.Errorf("Retrieved car mark mismatch: expected Honda, got %s", retrievedCar.Mark)
	}

	err = carsRepo.DeleteCar(car.RegNumber)
	if err != nil {
		t.Errorf("Failed to delete car: %v", err)
	}
	//
	err = ownersRepo.DeleteOwner(owner.ID)
	if err != nil {
		t.Errorf("Failed to delete owner: %v", err)
	}
}
