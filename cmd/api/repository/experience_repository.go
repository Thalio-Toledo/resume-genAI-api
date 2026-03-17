package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/cmd/api/model"
)

type ExperienceRepository struct {
	experiences []model.Experience
	db          *sql.DB
}

func NewExperienceRepository(db *sql.DB) *ExperienceRepository {
	return &ExperienceRepository{
		experiences: []model.Experience{},
		db:          db,
	}
}

func (r *ExperienceRepository) Get() ([]model.Experience, error) {
	query := `
		SELECT
			experience_id,
			profile_id,
			company,
			is_current,
			role,
			description,
			start_date,
			end_date
		FROM experience
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []model.Experience

	for rows.Next() {
		var experience model.Experience

		err := rows.Scan(
			&experience.ExperienceId,
			&experience.ProfileID,
			&experience.Company,
			&experience.IsCurrent,
			&experience.Role,
			&experience.Description,
			&experience.StartDate,
			&experience.EndDate,
		)
		if err != nil {
			return nil, err
		}

		experiences = append(experiences, experience)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return experiences, nil
}

func (r *ExperienceRepository) FindByID(id string) (*model.Experience, error) {
	query := `
		SELECT
			experience_id,
			profile_id,
			company,
			is_current,
			role,
			description,
			start_date,
			end_date
		FROM experience
		WHERE experience_id = @id
	`
	var experience model.Experience

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&experience.ExperienceId,
		&experience.ProfileID,
		&experience.Company,
		&experience.IsCurrent,
		&experience.Role,
		&experience.Description,
		&experience.StartDate,
		&experience.EndDate,
	)
	if err != nil {
		return nil, err
	}

	return &experience, nil
}

func (r *ExperienceRepository) FindByProfileID(profileID int) ([]model.Experience, error) {
	query := `
		SELECT
			experience_id,
			profile_id,
			company,
			is_current,
			role,
			description,
			start_date,
			end_date
		FROM experience
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []model.Experience

	for rows.Next() {
		var experience model.Experience

		err := rows.Scan(
			&experience.ExperienceId,
			&experience.ProfileID,
			&experience.Company,
			&experience.IsCurrent,
			&experience.Role,
			&experience.Description,
			&experience.StartDate,
			&experience.EndDate,
		)
		if err != nil {
			return nil, err
		}

		experiences = append(experiences, experience)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return experiences, nil
}

func (r *ExperienceRepository) Create(experience model.Experience) (string, error) {
	query := `
		INSERT INTO experience (
			profile_id,
			company,
			is_current,
			role,
			description,
			start_date,
			end_date
		)
		OUTPUT INSERTED.experience_id
		VALUES (
			@profile_id,
			@company,
			@is_current,
			@role,
			@description,
			@start_date,
			@end_date
		)
	`

	var id string

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", experience.ProfileID),
		sql.Named("company", experience.Company),
		sql.Named("is_current", experience.IsCurrent),
		sql.Named("role", experience.Role),
		sql.Named("description", experience.Description),
		sql.Named("start_date", experience.StartDate),
		sql.Named("end_date", experience.EndDate),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ExperienceRepository) Update(experience model.Experience) (bool, error) {
	query := `
		UPDATE experience
		SET
			company = @company,
			is_current = @is_current,
			role = @role,
			description = @description,
			start_date = @start_date,
			end_date = @end_date
		WHERE experience_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", experience.ExperienceId),
		sql.Named("company", experience.Company),
		sql.Named("is_current", experience.IsCurrent),
		sql.Named("role", experience.Role),
		sql.Named("description", experience.Description),
		sql.Named("start_date", experience.StartDate),
		sql.Named("end_date", experience.EndDate),
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

func (r *ExperienceRepository) Delete(id string) (bool, error) {
	query := `
		DELETE experience
		WHERE experience_id = @id
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
