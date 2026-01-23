package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type SkillUseCase struct {
	repo *repository.SkillRepository
}

func NewSkillUseCase(repo *repository.SkillRepository) *SkillUseCase {
	return &SkillUseCase{repo: repo}
}

func (uc *SkillUseCase) GetAll() []model.Skill {
	return uc.repo.GetAll()
}

func (uc *SkillUseCase) FindByID(id string) (*model.Skill, error) {
	return uc.repo.FindByID(id)
}

func (uc *SkillUseCase) Create(skill model.Skill) (string, error) {
	return uc.repo.Create(skill)
}

func (uc *SkillUseCase) Update(skill model.Skill) (bool, error) {
	return uc.repo.Update(skill)
}

func (uc *SkillUseCase) Delete(id string) (bool, error) {
	return uc.repo.Delete(id)
}
