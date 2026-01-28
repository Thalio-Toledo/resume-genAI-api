package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type SkillRepository struct {
	skills []model.Skill
}

func NewSkillRepository() *SkillRepository {
	return &SkillRepository{
		skills: []model.Skill{},
	}
}

func (r *SkillRepository) GetAll() []model.Skill {
	return r.skills
}

func (r *SkillRepository) FindByID(id string) (*model.Skill, error) {
	for _, s := range r.skills {
		if s.SkillId == id {
			return &s, nil
		}
	}
	return nil, errors.New("Skill not found")
}

func (r *SkillRepository) Create(skill model.Skill) (string, error) {
	r.skills = append(r.skills, skill)
	return skill.SkillId, nil
}

func (r *SkillRepository) Update(skill model.Skill) (bool, error) {
	for i, s := range r.skills {
		if s.SkillId == skill.SkillId {
			r.skills[i] = skill
			return true, nil
		}
	}
	return false, errors.New("Skill not found")
}

func (r *SkillRepository) Delete(id string) (bool, error) {
	for i, s := range r.skills {
		if s.SkillId == id {
			r.skills = append(r.skills[:i], r.skills[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Skill not found")
}
