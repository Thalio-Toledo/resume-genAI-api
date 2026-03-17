package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/cmd/api/model"
)

type EducationRepository struct {
	educations []model.Education
	db         *sql.DB
}

func NewEducationRepository(db *sql.DB) *EducationRepository {
	return &EducationRepository{
		educations: []model.Education{},
		db:         db,
	}
}

func (r *EducationRepository) Get() ([]model.Education, error) {
	query := `
		SELECT
			education_id,
			profile_id,
			institution,
			degree,
			field,
			start_date,
			end_date
		FROM education
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var educations []model.Education

	for rows.Next() {
		var education model.Education

		err := rows.Scan(
			&education.EducationId,
			&education.ProfileID,
			&education.Institution,
			&education.Degree,
			&education.Field,
			&education.StartDate,
			&education.EndDate,
		)
		if err != nil {
			return nil, err
		}

		educations = append(educations, education)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return educations, nil
}

func (r *EducationRepository) FindByID(id string) (*model.Education, error) {
	query := `
		SELECT
			education_id,
			profile_id,
			institution,
			degree,
			field,
			start_date,
			end_date
		FROM education
		WHERE education_id = @id
	`
	var education model.Education

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&education.EducationId,
		&education.ProfileID,
		&education.Institution,
		&education.Degree,
		&education.Field,
		&education.StartDate,
		&education.EndDate,
	)
	if err != nil {
		return nil, err
	}

	return &education, nil
}

func (r *EducationRepository) FindByProfileID(profileID int) ([]model.Education, error) {
	query := `
		SELECT
			education_id,
			profile_id,
			institution,
			degree,
			field,
			start_date,
			end_date
		FROM education
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var educations []model.Education

	for rows.Next() {
		var education model.Education

		err := rows.Scan(
			&education.EducationId,
			&education.ProfileID,
			&education.Institution,
			&education.Degree,
			&education.Field,
			&education.StartDate,
			&education.EndDate,
		)
		if err != nil {
			return nil, err
		}

		educations = append(educations, education)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return educations, nil
}

func (r *EducationRepository) Create(education model.Education) (string, error) {
	query := `
		INSERT INTO education (
			profile_id,
			institution,
			degree,
			field,
			start_date,
			end_date
		)
		OUTPUT INSERTED.education_id
		VALUES (
			@profile_id,
			@institution,
			@degree,
			@field,
			@start_date,
			@end_date
		)
	`

	var id string

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", education.ProfileID),
		sql.Named("institution", education.Institution),
		sql.Named("degree", education.Degree),
		sql.Named("field", education.Field),
		sql.Named("start_date", education.StartDate),
		sql.Named("end_date", education.EndDate),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *EducationRepository) Update(education model.Education) (bool, error) {
	query := `
		UPDATE education
		SET
			institution = @institution,
			degree = @degree,
			field = @field,
			start_date = @start_date,
			end_date = @end_date
		WHERE education_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", education.EducationId),
		sql.Named("institution", education.Institution),
		sql.Named("degree", education.Degree),
		sql.Named("field", education.Field),
		sql.Named("start_date", education.StartDate),
		sql.Named("end_date", education.EndDate),
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

func (r *EducationRepository) Delete(id string) (bool, error) {
	query := `
		DELETE education
		WHERE education_id = @id
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
