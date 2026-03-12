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
 			FROM project
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
			&project.ProfileID,
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
			,profile_id
			,name
			,description
			,link
			,active
		FROM project 
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
		FROM project 
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
	query := `
		INSERT INTO project (
			profile_id,
			name,
			description,
			link,
			active
		)
		OUTPUT INSERTED.project_id
		VALUES (
			@profile_id,
			@name,
			@description,
			@link,
			@active
		)
	`

	var id string

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", proj.ProfileID),
		sql.Named("name", proj.Name),
		sql.Named("description", proj.Description),
		sql.Named("link", proj.Link),
		sql.Named("active", proj.Active),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ProjectRepository) Update(proj model.Project) (bool, error) {
	query := `
		UPDATE project
		SET
			name = @name,
			description = @description,
			link = @link,
			active = @active
		WHERE project_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", proj.ProjectId),
		sql.Named("name", proj.Name),
		sql.Named("description", proj.Description),
		sql.Named("link", proj.Link),
		sql.Named("active", proj.Active),
	)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, sql.ErrNoRows
	}

	return true, nil
}

func (r *ProjectRepository) Delete(id string) (bool, error) {
	query := `
		DELETE project
		WHERE project_id = @id
	`
	result, err := r.db.Exec(query, sql.Named("id", id))

	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, errors.New("no rows affected")
	}
	return true, nil
}
