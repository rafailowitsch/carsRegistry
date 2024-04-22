package repository

import (
	"carRegistry/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestCreateAndUpdateCar(t *testing.T) {
	owner := domain.Owners{
		ID:      uuid.New(),
		Name:    "Bob",
		Surname: "Jones",
	}

	repo := NewRepository(db)

	_ = repo.OwnersRepo.CreateOwner(&owner)

	car := domain.Cars{
		RegNumber: "X999XX99",
		Mark:      "Toyota",
		Model:     "Camry",
		Year:      "2020",
		OwnerID:   owner.ID,
	}
	err := repo.CarsRepo.CreateCar(&car)
	if err != nil {
		t.Fatalf("Failed to create car: %v", err)
	}

	car.Mark = "Honda"
	err = repo.CarsRepo.UpdateCar(&car)
	if err != nil {
		t.Errorf("Failed to update car: %v", err)
	}

	retrievedCar, err := repo.CarsRepo.GetCarByRegNumber(car.RegNumber)
	if err != nil {
		t.Errorf("Failed to retrieve car: %v", err)
	}

	if retrievedCar.Mark != "Honda" {
		t.Errorf("Retrieved car mark mismatch: expected Honda, got %s", retrievedCar.Mark)
	}

	_ = repo.CarsRepo.DeleteCar(car.RegNumber) // Очистка после теста
	_ = repo.OwnersRepo.DeleteOwner(owner.ID)  // Удаление владельца
}
