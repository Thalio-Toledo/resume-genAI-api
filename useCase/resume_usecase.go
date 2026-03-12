package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type ResumeUseCase struct {
	repo *repository.ResumeRepository
}

func NewResumeUseCase(repo *repository.ResumeRepository) *ResumeUseCase {
	return &ResumeUseCase{repo: repo}
}

func (uc *ResumeUseCase) Get() ([]model.Resume, error) {
	return uc.repo.Get()
}

func (uc *ResumeUseCase) FindByID(id string) (*model.Resume, error) {
	return uc.repo.FindByID(id)
}

func (uc *ResumeUseCase) FindByProfileID(profileID int) ([]model.Resume, error) {
	return uc.repo.FindByProfileID(profileID)
}

func (uc *ResumeUseCase) Create(resume model.Resume) (string, error) {
	return uc.repo.Create(resume)
}

func (uc *ResumeUseCase) Update(resume model.Resume) (bool, error) {
	return uc.repo.Update(resume)
}

func (uc *ResumeUseCase) Delete(id string) (bool, error) {
	return uc.repo.Delete(id)
}
