package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type ResumeRepository struct {
	resumes []model.Resume
}

func NewResumeRepository() *ResumeRepository {
	return &ResumeRepository{
		resumes: []model.Resume{},
	}
}

func (r *ResumeRepository) GetAll() []model.Resume {
	return r.resumes
}

func (r *ResumeRepository) FindByID(id string) (*model.Resume, error) {
	for _, res := range r.resumes {
		if res.ID == id {
			return &res, nil
		}
	}
	return nil, errors.New("Resume not found")
}

func (r *ResumeRepository) Create(resume model.Resume) (string, error) {
	r.resumes = append(r.resumes, resume)
	return resume.ID, nil
}

func (r *ResumeRepository) Update(resume model.Resume) (bool, error) {
	for i, res := range r.resumes {
		if res.ID == resume.ID {
			r.resumes[i] = resume
			return true, nil
		}
	}
	return false, errors.New("Resume not found")
}

func (r *ResumeRepository) Delete(id string) (bool, error) {
	for i, res := range r.resumes {
		if res.ID == id {
			r.resumes = append(r.resumes[:i], r.resumes[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Resume not found")
}
