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

func (uc *SocialMediaUseCase) Get() ([]model.SocialMedia, error) {
	return uc.repo.Get()
}

func (uc *SocialMediaUseCase) FindByID(id int) (*model.SocialMedia, error) {
	return uc.repo.FindByID(id)
}

func (uc *SocialMediaUseCase) FindByProfileID(profileID int) ([]model.SocialMedia, error) {
	return uc.repo.FindByProfileID(profileID)
}

func (uc *SocialMediaUseCase) Create(sm model.SocialMedia) (int, error) {
	return uc.repo.Create(sm)
}

func (uc *SocialMediaUseCase) Update(sm model.SocialMedia) (bool, error) {
	return uc.repo.Update(sm)
}

func (uc *SocialMediaUseCase) Delete(id int) (bool, error) {
	return uc.repo.Delete(id)
}
