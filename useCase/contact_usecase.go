package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type ContactUseCase struct {
	repo *repository.ContactRepository
}

func NewContactUseCase(repo *repository.ContactRepository) *ContactUseCase {
	return &ContactUseCase{repo: repo}
}

func (uc *ContactUseCase) GetAll() []model.Contact {
	return uc.repo.GetAll()
}

func (uc *ContactUseCase) FindByEmail(email string) (*model.Contact, error) {
	return uc.repo.FindByEmail(email)
}

func (uc *ContactUseCase) Create(contact model.Contact) (string, error) {
	return uc.repo.Create(contact)
}

func (uc *ContactUseCase) Update(contact model.Contact) (bool, error) {
	return uc.repo.Update(contact)
}

func (uc *ContactUseCase) Delete(email string) (bool, error) {
	return uc.repo.Delete(email)
}
