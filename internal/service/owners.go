package service

import (
	"carsRegistry/internal/domain"
	"carsRegistry/internal/repository"
	"github.com/google/uuid"
)

type OwnersService struct {
	OwnersRepo repository.Owners
}

func NewOwnersService(ownersRepo repository.Owners) *OwnersService {
	return &OwnersService{OwnersRepo: ownersRepo}
}

func (s *OwnersService) CreateOwner(ownerInput domain.OwnersInput) error {
	owner := &domain.Owners{
		ID:         uuid.New(),
		Name:       ownerInput.Name,
		Surname:    ownerInput.Surname,
		Patronymic: ownerInput.Patronymic,
	}
	return s.OwnersRepo.CreateOwner(owner)
}

func (s *OwnersService) GetOwnerByID(id string) (*domain.Owners, error) {
	ownerID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.OwnersRepo.GetOwnerByID(ownerID)
}
