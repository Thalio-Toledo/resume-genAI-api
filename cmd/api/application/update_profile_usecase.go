package application

import (
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/infrastructure"
)

type UpdateProfileUseCase struct {
	repo *infrastructure.ProfileRepository
}

func NewUpdateProfileUseCase(repo *infrastructure.ProfileRepository) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{repo: repo}
}

func (uc *ProfileUseCase) Update(profile domain.Profile) (bool, error) {
	return uc.repo.Update(profile)
}
