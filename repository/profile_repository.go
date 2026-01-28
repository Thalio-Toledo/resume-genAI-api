package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/model"
)

type ProfileRepository struct {
	profiles []model.Profile
	db       *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{
		profiles: []model.Profile{},
		db:       db,
	}
}

func (p *ProfileRepository) Get() ([]model.Profile, error) {
	query := `
		SELECT
			profile_id,
			name,
			email,
			birth_date,
			phone_number,
			description
		FROM profile
	`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []model.Profile

	for rows.Next() {
		var profile model.Profile

		err := rows.Scan(
			&profile.ProfileId,
			&profile.Name,
			&profile.Email,
			&profile.Birthdate,
			&profile.PhoneNumber,
			&profile.Description,
		)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (p *ProfileRepository) FindByID(id int) (*model.Profile, error) {
	query := `
		SELECT
			profile_id,
			name,
			email,
			birth_date,
			phone_number,
			description
		FROM profile
		WHERE profile_id = @id
	`

	var profile model.Profile

	err := p.db.QueryRow(query, sql.Named("id", id)).Scan(
		&profile.ProfileId,
		&profile.Name,
		&profile.Email,
		&profile.Birthdate,
		&profile.PhoneNumber,
		&profile.Description,
	)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (p *ProfileRepository) LoadProjects(profile *model.Profile) error {
	query := `
		SELECT
			 project_id
			,name
			,description
			,link
			,active
		FROM project 
		WHERE profile_id = @profile_id
	`
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
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
			return err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.Projects = projects
	return nil
}

// func (p *ProfileRepository) LoadCertifications(profile *model.Profile ) error{
// 	query := `

// 	`
// }

func (p *ProfileRepository) Create(profile model.Profile) (int, error) {
	query := `
		INSERT INTO profile (
			name,
			email,
			birth_date,
			phone_number,
			description
		)
		OUTPUT INSERTED.profile_id
		VALUES (
			@name,
			@email,
			@birth_date,
			@phone_number,
			@description
		)
	`

	var id int

	err := p.db.QueryRow(
		query,
		sql.Named("name", profile.Name),
		sql.Named("email", profile.Email),
		sql.Named("birth_date", profile.Birthdate),
		sql.Named("phone_number", profile.PhoneNumber),
		sql.Named("description", profile.Description),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *ProfileRepository) Update(profile model.Profile) (bool, error) {
	query := `
		UPDATE profile
		SET
			name = @name,
			email = @email,
			birth_date = @birth_date,
			phone_number = @phone_number,
			description = @description
		WHERE profile_id = @id
	`

	result, err := p.db.Exec(
		query,
		sql.Named("id", profile.ProfileId),
		sql.Named("name", profile.Name),
		sql.Named("email", profile.Email),
		sql.Named("birth_date", profile.Birthdate),
		sql.Named("phone_number", profile.PhoneNumber),
		sql.Named("description", profile.Description),
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

func (p *ProfileRepository) Delete(id int) (bool, error) {
	query := `
		DELETE profile
		WHERE profile_id = @id
	`
	result, err := p.db.Exec(query, sql.Named("id", id))

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
