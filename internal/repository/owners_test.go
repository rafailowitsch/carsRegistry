package repository

import (
	"carsRegistry/internal/domain"
	"github.com/google/uuid"
	"testing"
)

func TestCreateAndDeletePeople(t *testing.T) {
	owner := domain.Owners{
		ID:         uuid.New(),
		Name:       "Alice",
		Surname:    "Smith",
		Patronymic: "Middle",
	}

	ownersRepo := NewOwnersRepo(db)

	err := ownersRepo.CreateOwner(&owner)
	if err != nil {
		t.Errorf("Failed to create person: %v", err)
	}

	retrievedPerson, err := ownersRepo.GetOwnerByID(owner.ID)
	if err != nil {
		t.Errorf("Failed to retrieve person: %v", err)
	}

	if retrievedPerson.Name != "Alice" {
		t.Errorf("Retrieved person name mismatch: expected Alice, got %s", retrievedPerson.Name)
	}

	err = ownersRepo.DeleteOwner(owner.ID)
	if err != nil {
		t.Errorf("Failed to delete person: %v", err)
	}
}
