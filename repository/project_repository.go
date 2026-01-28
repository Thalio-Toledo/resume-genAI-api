package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/model"
)

type ProjectRepository struct {
	projects []model.Project
	db       *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		projects: []model.Project{},
		db:       db,
	}
}

func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	query := `
			SELECT
	   			 project_id
				,profile_id
				,name
				,description
				,link
				,active
 			FROM projects
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project

	for rows.Next() {
		var project model.Project

		err := rows.Scan(
			&project.ProjectId,
			&project.Name,
			&project.Description,
			&project.Link,
			&project.Active,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return projects, err

}

func (r *ProjectRepository) FindByProfileID(profileID string) ([]model.Project, error) {
	query := `
		SELECT
			 project_id
			,name
			,description
			,link
			,active
		FROM projects 
		WHERE profile_id = @profile_id
	`

	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project

	for rows.Next() {
		var project model.Project

		err := rows.Scan(
			&project.ProjectId,
			&project.Name,
			&project.Description,
			&project.Link,
			&project.Active,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return projects, err
}

func (r *ProjectRepository) FindByID(id string) (*model.Project, error) {
	query := `
		SELECT
			 project_id
			,profile_id
			,name
			,description
			,link
			,active
		FROM projects 
		WHERE project_id = @id
	`

	var project model.Project

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&project.ProjectId,
		&project.Name,
		&project.Description,
		&project.Link,
		&project.Active,
	)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) Create(proj model.Project) (string, error) {
	r.projects = append(r.projects, proj)
	return proj.ProjectId, nil
}

func (r *ProjectRepository) Update(proj model.Project) (bool, error) {
	for i, p := range r.projects {
		if p.ProjectId == proj.ProjectId {
			r.projects[i] = proj
			return true, nil
		}
	}
	return false, errors.New("Project not found")
}

func (r *ProjectRepository) Delete(id string) (bool, error) {
	for i, p := range r.projects {
		if p.ProjectId == id {
			r.projects = append(r.projects[:i], r.projects[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Project not found")
}
