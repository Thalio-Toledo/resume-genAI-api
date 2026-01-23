package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type EducationUseCase struct {
	repo *repository.EducationRepository
}

func NewEducationUseCase(repo *repository.EducationRepository) *EducationUseCase {
	return &EducationUseCase{repo: repo}
}

func (uc *EducationUseCase) GetAll() []model.Education {
	return uc.repo.GetAll()
}

func (uc *EducationUseCase) FindByID(id string) (*model.Education, error) {
	return uc.repo.FindByID(id)
}

func (uc *EducationUseCase) Create(edu model.Education) (string, error) {
	return uc.repo.Create(edu)
}

func (uc *EducationUseCase) Update(edu model.Education) (bool, error) {
	return uc.repo.Update(edu)
}

func (uc *EducationUseCase) Delete(id string) (bool, error) {
	return uc.repo.Delete(id)
}
