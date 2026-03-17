package useCase

import (
	"resume-genAI-api/cmd/api/model"
	"resume-genAI-api/cmd/api/repository"
)

type ProjectUseCase struct {
	repo *repository.ProjectRepository
}

func NewProjectUseCase(repo *repository.ProjectRepository) *ProjectUseCase {
	return &ProjectUseCase{repo: repo}
}

func (uc *ProjectUseCase) Get() ([]model.Project, error) {
	return uc.repo.GetAll()
}

func (uc *ProjectUseCase) FindByID(id string) (*model.Project, error) {
	return uc.repo.FindByID(id)
}

func (uc *ProjectUseCase) FindByProfileID(profileID string) ([]model.Project, error) {
	return uc.repo.FindByProfileID(profileID)
}

func (uc *ProjectUseCase) Create(proj model.Project) (string, error) {
	return uc.repo.Create(proj)
}

func (uc *ProjectUseCase) Update(proj model.Project) (bool, error) {
	return uc.repo.Update(proj)
}

func (uc *ProjectUseCase) Delete(id string) (bool, error) {
	return uc.repo.Delete(id)
}
