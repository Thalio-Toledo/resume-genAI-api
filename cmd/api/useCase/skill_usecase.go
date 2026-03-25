package useCase

import (
	ai "resume-genAI-api/cmd/api/AI"
	"resume-genAI-api/cmd/api/model"
	"resume-genAI-api/cmd/api/repository"
)

type SkillUseCase struct {
	repo *repository.SkillRepository
}

func NewSkillUseCase(repo *repository.SkillRepository) *SkillUseCase {
	return &SkillUseCase{repo: repo}
}

func (uc *SkillUseCase) Get() ([]model.Skill, error) {
	return uc.repo.Get()
}

func (uc *SkillUseCase) FindByID(id string) (*model.Skill, error) {
	return uc.repo.FindByID(id)
}

func (uc *SkillUseCase) FindByProfileID(profileID int) ([]model.Skill, error) {
	return uc.repo.FindByProfileID(profileID)
}

func (uc *SkillUseCase) Create(skill *model.Skill) (int, error) {
	embeddings, _ := ai.GenerateEmbedding(skill.Name)
	return uc.repo.Create(skill, embeddings)
}

func (uc *SkillUseCase) Update(skill model.Skill) (bool, error) {
	return uc.repo.Update(skill)
}

func (uc *SkillUseCase) Delete(id string) (bool, error) {
	return uc.repo.Delete(id)
}
