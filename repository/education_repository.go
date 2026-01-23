package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type EducationRepository struct {
	educations []model.Education
}

func NewEducationRepository() *EducationRepository {
	return &EducationRepository{
		educations: []model.Education{},
	}
}

func (r *EducationRepository) GetAll() []model.Education {
	return r.educations
}

func (r *EducationRepository) FindByID(id string) (*model.Education, error) {
	for _, e := range r.educations {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, errors.New("Education not found")
}

func (r *EducationRepository) Create(edu model.Education) (string, error) {
	r.educations = append(r.educations, edu)
	return edu.ID, nil
}

func (r *EducationRepository) Update(edu model.Education) (bool, error) {
	for i, e := range r.educations {
		if e.ID == edu.ID {
			r.educations[i] = edu
			return true, nil
		}
	}
	return false, errors.New("Education not found")
}

func (r *EducationRepository) Delete(id string) (bool, error) {
	for i, e := range r.educations {
		if e.ID == id {
			r.educations = append(r.educations[:i], r.educations[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Education not found")
}
