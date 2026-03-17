package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"resume-genAI-api/cmd/api/model"
)

type SkillRepository struct {
	skills []model.Skill
	db     *sql.DB
}

func NewSkillRepository(db *sql.DB) *SkillRepository {
	return &SkillRepository{
		skills: []model.Skill{},
		db:     db,
	}
}

func (r *SkillRepository) Get() ([]model.Skill, error) {
	query := `
		SELECT
			skill_id,
			profile_id,
			name,
			level,
			embeddingsJSON
		FROM skill
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []model.Skill

	for rows.Next() {
		var skill model.Skill

		err := rows.Scan(
			&skill.SkillId,
			&skill.ProfileID,
			&skill.Name,
			&skill.Level,
			&skill.EmbeddingsJSON,
		)
		if err != nil {
			return nil, err
		}

		if skill.EmbeddingsJSON.Valid {
			err = json.Unmarshal([]byte(skill.EmbeddingsJSON.String), &skill.Embeddings)
			if err != nil {
				return nil, err
			}
		}

		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return skills, nil
}

func (r *SkillRepository) FindByID(id string) (*model.Skill, error) {
	query := `
		SELECT
			skill_id,
			profile_id,
			name,
			level,
			embeddingsJSON
		FROM skill
		WHERE skill_id = @id
	`
	var skill model.Skill

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&skill.SkillId,
		&skill.ProfileID,
		&skill.Name,
		&skill.Level,
		&skill.EmbeddingsJSON,
	)
	if err != nil {
		return nil, err
	}

	if skill.EmbeddingsJSON.Valid {
		err = json.Unmarshal([]byte(skill.EmbeddingsJSON.String), &skill.Embeddings)
		if err != nil {
			return nil, err
		}
	}

	return &skill, nil
}

func (r *SkillRepository) FindByProfileID(profileID int) ([]model.Skill, error) {
	query := `
		SELECT
			skill_id,
			profile_id,
			name,
			level,
			embeddingsJSON
		FROM skill
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []model.Skill

	for rows.Next() {
		var skill model.Skill

		err := rows.Scan(
			&skill.SkillId,
			&skill.ProfileID,
			&skill.Name,
			&skill.Level,
			&skill.EmbeddingsJSON,
		)
		if err != nil {
			return nil, err
		}

		if skill.EmbeddingsJSON.Valid {
			err = json.Unmarshal([]byte(skill.EmbeddingsJSON.String), &skill.Embeddings)
			if err != nil {
				return nil, err
			}
		}

		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return skills, nil
}

func (r *SkillRepository) Create(skill model.Skill, embeddings []float32) (string, error) {
	embeddingJSON, _ := json.Marshal(embeddings)
	query := `
		INSERT INTO skill (
			profile_id,
			name,
			level,
			embeddingsJSON
		)
		OUTPUT INSERTED.skill_id
		VALUES (
			@profile_id,
			@name,
			@level,
			@embeddingsJSON
		)
	`

	var id string

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", skill.ProfileID),
		sql.Named("name", skill.Name),
		sql.Named("level", skill.Level),
		sql.Named("embeddingsJSON", string(embeddingJSON)),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *SkillRepository) Update(skill model.Skill) (bool, error) {
	query := `
		UPDATE skill
		SET
			name = @name,
			level = @level
		WHERE skill_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", skill.SkillId),
		sql.Named("name", skill.Name),
		sql.Named("level", skill.Level),
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

func (r *SkillRepository) Delete(id string) (bool, error) {
	query := `
		DELETE skill
		WHERE skill_id = @id
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
