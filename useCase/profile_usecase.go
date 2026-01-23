package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type ProfileUseCase struct {
	repo *repository.ProfileRepository
}

func NewProfileUseCase(repo *repository.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{repo: repo}
}

func (uc *ProfileUseCase) Get() ([]model.Profile, error) {
	return uc.repo.Get()
}

func (uc *ProfileUseCase) FindByID(id int) (*model.Profile, error) {
	return uc.repo.FindByID(id)
}

func (uc *ProfileUseCase) Create(profile model.Profile) (int, error) {
	return uc.repo.Create(profile)
}

func (uc *ProfileUseCase) Update(profile model.Profile) (bool, error) {
	return uc.repo.Update(profile)
}

func (uc *ProfileUseCase) Delete(id int) (bool, error) {
	return uc.repo.Delete(id)
}
