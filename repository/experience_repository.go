
package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type ExperienceRepository struct {
	experiences []model.Experience
}

func NewExperienceRepository() *ExperienceRepository {
	return &ExperienceRepository{
		experiences: []model.Experience{},
	}
}

func (r *ExperienceRepository) GetAll() []model.Experience {
	return r.experiences
}

func (r *ExperienceRepository) FindByID(id string) (*model.Experience, error) {
	for _, e := range r.experiences {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, errors.New("Experience not found")
}

func (r *ExperienceRepository) Create(exp model.Experience) (string, error) {
	r.experiences = append(r.experiences, exp)
	return exp.ID, nil
}

func (r *ExperienceRepository) Update(exp model.Experience) (bool, error) {
	for i, e := range r.experiences {
		if e.ID == exp.ID {
			r.experiences[i] = exp
			return true, nil
		}
	}
	return false, errors.New("Experience not found")
}

func (r *ExperienceRepository) Delete(id string) (bool, error) {
	for i, e := range r.experiences {
		if e.ID == id {
			r.experiences = append(r.experiences[:i], r.experiences[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Experience not found")
}
