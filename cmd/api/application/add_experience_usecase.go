package application

import (
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/infrastructure"
)

type AddExperienceUseCase struct {
	repo               *infrastructure.ProfileRepository
	findProfileUseCase *FindProfileUseCase
}

func NewAddExperienceUseCase(repo *infrastructure.ProfileRepository, findProfileUseCase *FindProfileUseCase) *AddExperienceUseCase {
	return &AddExperienceUseCase{repo: repo, findProfileUseCase: findProfileUseCase}
}

func (uc *AddExperienceUseCase) AddExperience(experience domain.Experience) (int, error) {
	p, err := uc.findProfileUseCase.FindByID(experience.ProfileID)

	if err != nil {
		return 0, err
	}

	if err := p.AddExperience(experience); err != nil {
		return 0, err
	}

	return uc.repo.AddExperience(experience)
}
