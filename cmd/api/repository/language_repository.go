package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/cmd/api/model"
)

type LanguageRepository struct {
	languages []model.Language
	db        *sql.DB
}

func NewLanguageRepository(db *sql.DB) *LanguageRepository {
	return &LanguageRepository{
		languages: []model.Language{},
		db:        db,
	}
}

func (r *LanguageRepository) Get() ([]model.Language, error) {
	query := `
		SELECT
			language_id,
			profile_id,
			name,
			level
		FROM language
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []model.Language

	for rows.Next() {
		var language model.Language

		err := rows.Scan(
			&language.LanguageId,
			&language.ProfileID,
			&language.Name,
			&language.Level,
		)
		if err != nil {
			return nil, err
		}

		languages = append(languages, language)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return languages, nil
}

func (r *LanguageRepository) FindByID(id string) (*model.Language, error) {
	query := `
		SELECT
			language_id,
			profile_id,
			name,
			level
		FROM language
		WHERE language_id = @id
	`
	var language model.Language

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&language.LanguageId,
		&language.ProfileID,
		&language.Name,
		&language.Level,
	)
	if err != nil {
		return nil, err
	}

	return &language, nil
}

func (r *LanguageRepository) FindByProfileID(profileID int) ([]model.Language, error) {
	query := `
		SELECT
			language_id,
			profile_id,
			name,
			level
		FROM language
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []model.Language

	for rows.Next() {
		var language model.Language

		err := rows.Scan(
			&language.LanguageId,
			&language.ProfileID,
			&language.Name,
			&language.Level,
		)
		if err != nil {
			return nil, err
		}

		languages = append(languages, language)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return languages, nil
}

func (r *LanguageRepository) Create(language model.Language) (string, error) {
	query := `
		INSERT INTO language (
			profile_id,
			name,
			level
		)
		OUTPUT INSERTED.language_id
		VALUES (
			@profile_id,
			@name,
			@level
		)
	`

	var id string

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", language.ProfileID),
		sql.Named("name", language.Name),
		sql.Named("level", language.Level),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *LanguageRepository) Update(language model.Language) (bool, error) {
	query := `
		UPDATE language
		SET
			name = @name,
			level = @level
		WHERE language_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", language.LanguageId),
		sql.Named("name", language.Name),
		sql.Named("level", language.Level),
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

func (r *LanguageRepository) Delete(id string) (bool, error) {
	query := `
		DELETE language
		WHERE language_id = @id
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
