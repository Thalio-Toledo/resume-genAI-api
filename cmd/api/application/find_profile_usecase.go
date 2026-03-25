package application

import (
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/infrastructure"
)

type FindProfileUseCase struct {
	repo *infrastructure.ProfileRepository
}

func NewFindProfileUseCase(repo *infrastructure.ProfileRepository) *FindProfileUseCase {
	return &FindProfileUseCase{repo: repo}
}

func (uc *FindProfileUseCase) FindByID(id int) (*domain.Profile, error) {
	profile, _ := uc.repo.FindByID(id)
	uc.repo.LoadEducations(profile)
	uc.repo.LoadProjects(profile)
	uc.repo.LoadCertifications(profile)
	uc.repo.LoadExperiences(profile)
	uc.repo.LoadSkill(profile)
	uc.repo.LoadLanguages(profile)
	uc.repo.LoadSocialMedia(profile)

	return profile, nil
}
