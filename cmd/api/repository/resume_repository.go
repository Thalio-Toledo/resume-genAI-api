package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/cmd/api/model"
)

type ResumeRepository struct {
	resumes []model.Resume
	db      *sql.DB
}

func NewResumeRepository(db *sql.DB) *ResumeRepository {
	return &ResumeRepository{
		resumes: []model.Resume{},
		db:      db,
	}
}

func (r *ResumeRepository) Get() ([]model.Resume, error) {
	query := `
		SELECT
			id,
			profile_id,
			skill_ids,
			experience_ids,
			project_ids,
			certification_ids,
			education_ids,
			language_ids,
			social_media_ids,
			contact_ids
		FROM resume
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []model.Resume

	for rows.Next() {
		var resume model.Resume

		err := rows.Scan(
			&resume.ID,
			&resume.ProfileID,
			&resume.SkillIDs,
			&resume.ExperienceIDs,
			&resume.ProjectIDs,
			&resume.CertificationIDs,
			&resume.EducationIDs,
			&resume.LanguageIDs,
			&resume.SocialMediaIDs,
			&resume.ContactIDs,
		)
		if err != nil {
			return nil, err
		}

		resumes = append(resumes, resume)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resumes, nil
}

func (r *ResumeRepository) FindByID(id string) (*model.Resume, error) {
	query := `
		SELECT
			id,
			profile_id,
			skill_ids,
			experience_ids,
			project_ids,
			certification_ids,
			education_ids,
			language_ids,
			social_media_ids,
			contact_ids
		FROM resume
		WHERE id = @id
	`
	var resume model.Resume

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&resume.ID,
		&resume.ProfileID,
		&resume.SkillIDs,
		&resume.ExperienceIDs,
		&resume.ProjectIDs,
		&resume.CertificationIDs,
		&resume.EducationIDs,
		&resume.LanguageIDs,
		&resume.SocialMediaIDs,
		&resume.ContactIDs,
	)
	if err != nil {
		return nil, err
	}

	return &resume, nil
}

func (r *ResumeRepository) FindByProfileID(profileID int) ([]model.Resume, error) {
	query := `
		SELECT
			id,
			profile_id,
			skill_ids,
			experience_ids,
			project_ids,
			certification_ids,
			education_ids,
			language_ids,
			social_media_ids,
			contact_ids
		FROM resume
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []model.Resume

	for rows.Next() {
		var resume model.Resume

		err := rows.Scan(
			&resume.ID,
			&resume.ProfileID,
			&resume.SkillIDs,
			&resume.ExperienceIDs,
			&resume.ProjectIDs,
			&resume.CertificationIDs,
			&resume.EducationIDs,
			&resume.LanguageIDs,
			&resume.SocialMediaIDs,
			&resume.ContactIDs,
		)
		if err != nil {
			return nil, err
		}

		resumes = append(resumes, resume)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resumes, nil
}

func (r *ResumeRepository) Create(resume model.Resume) (string, error) {
	query := `
		INSERT INTO resume (
			profile_id,
			skill_ids,
			experience_ids,
			project_ids,
			certification_ids,
			education_ids,
			language_ids,
			social_media_ids,
			contact_ids
		)
		OUTPUT INSERTED.id
		VALUES (
			@profile_id,
			@skill_ids,
			@experience_ids,
			@project_ids,
			@certification_ids,
			@education_ids,
			@language_ids,
			@social_media_ids,
			@contact_ids
		)
	`

	var id string

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", resume.ProfileID),
		sql.Named("skill_ids", resume.SkillIDs),
		sql.Named("experience_ids", resume.ExperienceIDs),
		sql.Named("project_ids", resume.ProjectIDs),
		sql.Named("certification_ids", resume.CertificationIDs),
		sql.Named("education_ids", resume.EducationIDs),
		sql.Named("language_ids", resume.LanguageIDs),
		sql.Named("social_media_ids", resume.SocialMediaIDs),
		sql.Named("contact_ids", resume.ContactIDs),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ResumeRepository) Update(resume model.Resume) (bool, error) {
	query := `
		UPDATE resume
		SET
			skill_ids = @skill_ids,
			experience_ids = @experience_ids,
			project_ids = @project_ids,
			certification_ids = @certification_ids,
			education_ids = @education_ids,
			language_ids = @language_ids,
			social_media_ids = @social_media_ids,
			contact_ids = @contact_ids
		WHERE id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", resume.ID),
		sql.Named("skill_ids", resume.SkillIDs),
		sql.Named("experience_ids", resume.ExperienceIDs),
		sql.Named("project_ids", resume.ProjectIDs),
		sql.Named("certification_ids", resume.CertificationIDs),
		sql.Named("education_ids", resume.EducationIDs),
		sql.Named("language_ids", resume.LanguageIDs),
		sql.Named("social_media_ids", resume.SocialMediaIDs),
		sql.Named("contact_ids", resume.ContactIDs),
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

func (r *ResumeRepository) Delete(id string) (bool, error) {
	query := `
		DELETE resume
		WHERE id = @id
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
