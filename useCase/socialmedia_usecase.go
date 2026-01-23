package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type SocialMediaUseCase struct {
	repo *repository.SocialMediaRepository
}

func NewSocialMediaUseCase(repo *repository.SocialMediaRepository) *SocialMediaUseCase {
	return &SocialMediaUseCase{repo: repo}
}

func (uc *SocialMediaUseCase) GetAll() []model.SocialMedia {
	return uc.repo.GetAll()
}

func (uc *SocialMediaUseCase) FindByHandle(handle string) (*model.SocialMedia, error) {
	return uc.repo.FindByHandle(handle)
}

func (uc *SocialMediaUseCase) Create(sm model.SocialMedia) (string, error) {
	return uc.repo.Create(sm)
}

func (uc *SocialMediaUseCase) Update(sm model.SocialMedia) (bool, error) {
	return uc.repo.Update(sm)
}

func (uc *SocialMediaUseCase) Delete(handle string) (bool, error) {
	return uc.repo.Delete(handle)
}
