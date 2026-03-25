package application

import (
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/infrastructure"
)

type ProfileUseCase struct {
	repo *infrastructure.ProfileRepository
}

func NewProfileUseCase(repo *infrastructure.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{repo: repo}
}

func (uc *ProfileUseCase) Get() ([]domain.Profile, error) {
	return uc.repo.Get()
}
