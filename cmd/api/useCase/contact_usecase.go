package useCase

import (
	"resume-genAI-api/cmd/api/model"
	"resume-genAI-api/cmd/api/repository"
)

type ContactUseCase struct {
	repo *repository.ContactRepository
}

func NewContactUseCase(repo *repository.ContactRepository) *ContactUseCase {
	return &ContactUseCase{repo: repo}
}

func (uc *ContactUseCase) Get() ([]model.Contact, error) {
	return uc.repo.Get()
}

func (uc *ContactUseCase) FindByID(id int) (*model.Contact, error) {
	return uc.repo.FindByID(id)
}

func (uc *ContactUseCase) FindByProfileID(profileID int) ([]model.Contact, error) {
	return uc.repo.FindByProfileID(profileID)
}

func (uc *ContactUseCase) Create(contact model.Contact) (int, error) {
	return uc.repo.Create(contact)
}

func (uc *ContactUseCase) Update(contact model.Contact) (bool, error) {
	return uc.repo.Update(contact)
}

func (uc *ContactUseCase) Delete(id int) (bool, error) {
	return uc.repo.Delete(id)
}
