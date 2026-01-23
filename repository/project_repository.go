package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type ProjectRepository struct {
	projects []model.Project
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{
		projects: []model.Project{},
	}
}

func (r *ProjectRepository) GetAll() []model.Project {
	return r.projects
}

func (r *ProjectRepository) FindByID(id string) (*model.Project, error) {
	for _, p := range r.projects {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("Project not found")
}

func (r *ProjectRepository) Create(proj model.Project) (string, error) {
	r.projects = append(r.projects, proj)
	return proj.ID, nil
}

func (r *ProjectRepository) Update(proj model.Project) (bool, error) {
	for i, p := range r.projects {
		if p.ID == proj.ID {
			r.projects[i] = proj
			return true, nil
		}
	}
	return false, errors.New("Project not found")
}

func (r *ProjectRepository) Delete(id string) (bool, error) {
	for i, p := range r.projects {
		if p.ID == id {
			r.projects = append(r.projects[:i], r.projects[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Project not found")
}
