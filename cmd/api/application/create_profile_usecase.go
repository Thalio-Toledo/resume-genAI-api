package application

import (
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/infrastructure"
)

type CreateProfileUseCase struct {
	repo *infrastructure.ProfileRepository
}

func NewCreateProfileUseCase(repo *infrastructure.ProfileRepository) *CreateProfileUseCase {
	return &CreateProfileUseCase{repo: repo}
}

func (uc *CreateProfileUseCase) Create(profile domain.Profile) (int, error) {
	return uc.repo.Create(profile)
}
