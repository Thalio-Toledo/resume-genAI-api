package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type LanguageUseCase struct {
	repo *repository.LanguageRepository
}

func NewLanguageUseCase(repo *repository.LanguageRepository) *LanguageUseCase {
	return &LanguageUseCase{repo: repo}
}

func (uc *LanguageUseCase) GetAll() []model.Language {
	return uc.repo.GetAll()
}

func (uc *LanguageUseCase) FindByID(id string) (*model.Language, error) {
	return uc.repo.FindByID(id)
}

func (uc *LanguageUseCase) Create(lang model.Language) (string, error) {
	return uc.repo.Create(lang)
}

func (uc *LanguageUseCase) Update(lang model.Language) (bool, error) {
	return uc.repo.Update(lang)
}

func (uc *LanguageUseCase) Delete(id string) (bool, error) {
	return uc.repo.Delete(id)
}
