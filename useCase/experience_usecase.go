package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type ExperienceUseCase struct {
	repo *repository.ExperienceRepository
}

func NewExperienceUseCase(repo *repository.ExperienceRepository) *ExperienceUseCase {
	return &ExperienceUseCase{repo: repo}
}

func (uc *ExperienceUseCase) GetAll() []model.Experience {
	return uc.repo.GetAll()
}

func (uc *ExperienceUseCase) FindByID(id string) (*model.Experience, error) {
	return uc.repo.FindByID(id)
}

func (uc *ExperienceUseCase) Create(exp model.Experience) (string, error) {
	return uc.repo.Create(exp)
}

func (uc *ExperienceUseCase) Update(exp model.Experience) (bool, error) {
	return uc.repo.Update(exp)
}

func (uc *ExperienceUseCase) Delete(id string) (bool, error) {
	return uc.repo.Delete(id)
}
