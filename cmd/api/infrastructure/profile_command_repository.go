package infrastructure

import (
	"database/sql"
	"encoding/json"
	"errors"
	"resume-genAI-api/cmd/api/domain"
)

type ProfileCommandRepository struct {
	profiles []domain.Profile
	db       *sql.DB
}

func NewProfileCommandRepository(db *sql.DB) *ProfileCommandRepository {
	return &ProfileCommandRepository{
		profiles: []domain.Profile{},
		db:       db,
	}
}

func (p *ProfileCommandRepository) create(tx *sql.Tx, profile *domain.Profile) (int, error) {
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

	err := tx.QueryRow(
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

	profile.ProfileId = id

	return id, nil
}

func (p *ProfileCommandRepository) update(tx *sql.Tx, profile *domain.Profile) (bool, error) {
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

func (p *ProfileCommandRepository) Delete(id int) error {
	query := `
		DELETE profile
		WHERE profile_id = @id
	`
	result, err := p.db.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) Save(profile *domain.Profile) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if profile.ProfileId == 0 {
		if _, err := p.create(tx, profile); err != nil {
			return err
		}
	} else {
		if _, err := p.update(tx, profile); err != nil {
			return err
		}
	}

	if err := p.saveProjects(tx, profile); err != nil {
		return err
	}

	if err := p.saveEducations(tx, profile); err != nil {
		return err
	}

	if err := p.saveCertifications(tx, profile); err != nil {
		return err
	}

	if err := p.saveExperiences(tx, profile); err != nil {
		return err
	}

	if err := p.saveLanguages(tx, profile); err != nil {
		return err
	}

	if err := p.saveSocialMedias(tx, profile); err != nil {
		return err
	}

	if err := p.saveSkills(tx, profile); err != nil {
		return err
	}

	return nil
}

func (p *ProfileCommandRepository) loadProjects(tx *sql.Tx, profileId int) ([]domain.Project, error) {
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
	rows, err := tx.Query(query, sql.Named("profile_id", profileId))
	if err != nil {
		return []domain.Project{}, err
	}
	defer rows.Close()

	var projects []domain.Project

	for rows.Next() {
		var project domain.Project

		err := rows.Scan(
			&project.ProjectId,
			&project.Name,
			&project.Description,
			&project.Link,
			&project.Active,
		)
		if err != nil {
			return []domain.Project{}, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return []domain.Project{}, err
	}

	return projects, nil
}

func (p *ProfileCommandRepository) addProject(tx *sql.Tx, project *domain.Project) error {
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

	var id int

	err := tx.QueryRow(
		query,
		sql.Named("profile_id", project.ProfileId),
		sql.Named("name", project.Name),
		sql.Named("description", project.Description),
		sql.Named("link", project.Link),
		sql.Named("active", project.Active),
	).Scan(&id)

	if err != nil {
		return err
	}

	project.ProjectId = id

	return nil

}

func (p *ProfileCommandRepository) updateProject(tx *sql.Tx, project *domain.Project) error {
	query := `
		UPDATE project
		SET
			name = @name,
			description = @description,
			link = @link,
			active = @active
		WHERE project_id = @id
	`

	result, err := tx.Exec(
		query,
		sql.Named("id", project.ProjectId),
		sql.Named("name", project.Name),
		sql.Named("description", project.Description),
		sql.Named("link", project.Link),
		sql.Named("active", project.Active),
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *ProfileCommandRepository) deleteProject(tx *sql.Tx, id int) error {
	query := `
		DELETE project
		WHERE project_id = @id
	`
	result, err := p.db.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) loadCertifications(tx *sql.Tx, profileId int) ([]domain.Certification, error) {
	query := `
		SELECT 
			 certification_id
			,profile_id
			,name
			,issuer
			,date_issued
		FROM certification 
		WHERE profile_id = @profile_id
	`
	rows, err := tx.Query(query, sql.Named("profile_id", profileId))
	if err != nil {
		return []domain.Certification{}, err
	}
	defer rows.Close()

	var certifications []domain.Certification

	for rows.Next() {
		var certification domain.Certification

		err := rows.Scan(
			&certification.CertificationId,
			&certification.ProfileId,
			&certification.Name,
			&certification.Issuer,
			&certification.DateIssued,
		)
		if err != nil {
			return []domain.Certification{}, err
		}

		certifications = append(certifications, certification)

	}

	if err := rows.Err(); err != nil {
		return []domain.Certification{}, err
	}

	return certifications, err
}

func (p *ProfileCommandRepository) addCertification(tx *sql.Tx, certification *domain.Certification) error {
	query := `
		INSERT INTO certification (
			profile_id,
			name,
			issuer,
			date_issued
		)
		OUTPUT INSERTED.certification_id
		VALUES (
			@profile_id,
			@name,
			@issuer,
			@date_issued
		)
	`

	var id int

	err := tx.QueryRow(
		query,
		sql.Named("profile_id", certification.ProfileId),
		sql.Named("name", certification.Name),
		sql.Named("issuer", certification.Issuer),
		sql.Named("date_issued", certification.DateIssued),
	).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProfileCommandRepository) updateCertification(tx *sql.Tx, certification *domain.Certification) error {
	query := `
		UPDATE certification
		SET
			name = @name,
			issuer = @issuer,
			date_issued = @date_issued
		WHERE certification_id = @id
	`

	result, err := tx.Exec(
		query,
		sql.Named("id", certification.CertificationId),
		sql.Named("name", certification.Name),
		sql.Named("issuer", certification.Issuer),
		sql.Named("date_issued", certification.DateIssued),
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *ProfileCommandRepository) deleteCertification(tx *sql.Tx, id int) error {
	query := `
		DELETE certification
		WHERE certification_id = @id
	`
	result, err := tx.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) loadEducations(tx *sql.Tx, profileId int) ([]domain.Education, error) {
	query := `
		  SELECT 
			 education_id
			,profile_id
			,institution
			,degree
			,field
			,start_date
			,end_date
  		FROM education
		WHERE profile_id = @profile_id
	`
	rows, err := tx.Query(query, sql.Named("profile_id", profileId))
	if err != nil {
		return []domain.Education{}, err
	}
	defer rows.Close()

	var educations []domain.Education

	for rows.Next() {
		var education domain.Education

		err := rows.Scan(
			&education.EducationId,
			&education.ProfileId,
			&education.Institution,
			&education.Degree,
			&education.Field,
			&education.StartDate,
			&education.EndDate,
		)
		if err != nil {
			return []domain.Education{}, err
		}

		educations = append(educations, education)
	}

	if err := rows.Err(); err != nil {
		return []domain.Education{}, err
	}

	return educations, nil
}

func (p *ProfileCommandRepository) addEducation(tx *sql.Tx, education *domain.Education) error {
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

	var id int

	err := p.db.QueryRow(
		query,
		sql.Named("profile_id", education.ProfileId),
		sql.Named("institution", education.Institution),
		sql.Named("degree", education.Degree),
		sql.Named("field", education.Field),
		sql.Named("start_date", education.StartDate),
		sql.Named("end_date", education.EndDate),
	).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProfileCommandRepository) updateEducation(tx *sql.Tx, education *domain.Education) error {
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

	result, err := tx.Exec(
		query,
		sql.Named("id", education.EducationId),
		sql.Named("institution", education.Institution),
		sql.Named("degree", education.Degree),
		sql.Named("field", education.Field),
		sql.Named("start_date", education.StartDate),
		sql.Named("end_date", education.EndDate),
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *ProfileCommandRepository) deleteEducation(tx *sql.Tx, id int) error {
	query := `
		DELETE education
		WHERE education_id = @id
	`
	result, err := tx.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) loadExperiences(tx *sql.Tx, profileId int) ([]domain.Experience, error) {
	query := `
		  SELECT 
			 experience_id
			,profile_id
			,company
			,is_current
			,role
			,description
			,start_date
			,end_date
  		FROM experience
		WHERE profile_id = @profile_id
	`
	rows, err := tx.Query(query, sql.Named("profile_id", profileId))
	if err != nil {
		return []domain.Experience{}, err
	}
	defer rows.Close()

	var experiences []domain.Experience

	for rows.Next() {
		var experience domain.Experience

		err := rows.Scan(
			&experience.ExperienceId,
			&experience.ProfileId,
			&experience.Company,
			&experience.IsCurrent,
			&experience.Role,
			&experience.Description,
			&experience.StartDate,
			&experience.EndDate,
		)
		if err != nil {
			return []domain.Experience{}, err
		}

		experiences = append(experiences, experience)
	}

	if err := rows.Err(); err != nil {
		return []domain.Experience{}, err
	}

	return experiences, nil
}

func (p *ProfileCommandRepository) addExperience(tx *sql.Tx, experience *domain.Experience) error {
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

	var id int

	err := tx.QueryRow(
		query,
		sql.Named("profile_id", experience.ProfileId),
		sql.Named("company", experience.Company),
		sql.Named("is_current", experience.IsCurrent),
		sql.Named("role", experience.Role),
		sql.Named("description", experience.Description),
		sql.Named("start_date", experience.StartDate),
		sql.Named("end_date", experience.EndDate),
	).Scan(&id)

	if err != nil {
		return err
	}

	experience.ExperienceId = id

	return nil
}

func (p *ProfileCommandRepository) updateExperience(tx *sql.Tx, experience *domain.Experience) error {
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

	result, err := tx.Exec(
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
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *ProfileCommandRepository) deleteExperience(tx *sql.Tx, id int) error {
	query := `
		DELETE experience
		WHERE experience_id = @id
	`
	result, err := tx.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) loadLanguages(tx *sql.Tx, profileId int) ([]domain.Language, error) {
	query := `
		  SELECT 
			 language_id
			,profile_id
			,name
			,level
  		FROM language
		WHERE profile_id = @profile_id
	`
	rows, err := tx.Query(query, sql.Named("profile_id", profileId))
	if err != nil {
		return []domain.Language{}, err
	}
	defer rows.Close()

	var languages []domain.Language

	for rows.Next() {
		var language domain.Language

		err := rows.Scan(
			&language.LanguageId,
			&language.ProfileId,
			&language.Name,
			&language.Level,
		)
		if err != nil {
			return []domain.Language{}, err
		}

		languages = append(languages, language)
	}

	if err := rows.Err(); err != nil {
		return []domain.Language{}, err
	}

	return languages, nil
}

func (p *ProfileCommandRepository) addLanguage(tx *sql.Tx, language *domain.Language) error {
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

	var id int

	err := tx.QueryRow(
		query,
		sql.Named("profile_id", language.ProfileId),
		sql.Named("name", language.Name),
		sql.Named("level", language.Level),
	).Scan(&id)

	if err != nil {
		return err
	}

	language.LanguageId = id

	return nil
}

func (p *ProfileCommandRepository) updateLanguage(tx *sql.Tx, language *domain.Language) error {
	query := `
		UPDATE language
		SET
			name = @name,
			level = @level
		WHERE language_id = @id
	`

	result, err := tx.Exec(
		query,
		sql.Named("id", language.LanguageId),
		sql.Named("name", language.Name),
		sql.Named("level", language.Level),
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *ProfileCommandRepository) deleteLanguage(tx *sql.Tx, id int) error {
	query := `
		DELETE language
		WHERE language_id = @id
	`
	result, err := tx.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) loadSkills(tx *sql.Tx, profileId int) ([]domain.Skill, error) {
	query := `
		  SELECT 
			 skill_id
			,profile_id
			,name
			,level
			,embeddingsJSON
  		FROM skill
		WHERE profile_id = @profile_id
	`
	rows, err := tx.Query(query, sql.Named("profile_id", profileId))
	if err != nil {
		return []domain.Skill{}, err
	}
	defer rows.Close()

	var skills []domain.Skill

	for rows.Next() {
		var skill domain.Skill

		err := rows.Scan(
			&skill.SkillId,
			&skill.ProfileId,
			&skill.Name,
			&skill.Level,
			&skill.EmbeddingsJSON,
		)
		if err != nil {
			return []domain.Skill{}, err
		}

		if skill.EmbeddingsJSON.Valid {
			err = json.Unmarshal([]byte(skill.EmbeddingsJSON.String), &skill.Embeddings)
			if err != nil {
				return []domain.Skill{}, err
			}
		}

		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return []domain.Skill{}, err
	}

	return skills, nil
}

func (p *ProfileCommandRepository) addSkill(tx *sql.Tx, skill *domain.Skill) error {
	embeddingJSON, _ := json.Marshal(skill.Embeddings)
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

	var id int

	err := tx.QueryRow(
		query,
		sql.Named("profile_id", skill.ProfileId),
		sql.Named("name", skill.Name),
		sql.Named("level", skill.Level),
		sql.Named("embeddingsJSON", string(embeddingJSON)),
	).Scan(&id)

	if err != nil {
		return err
	}

	skill.SkillId = id

	return nil
}

func (p *ProfileCommandRepository) updateSkill(tx *sql.Tx, skill *domain.Skill) error {
	query := `
		UPDATE skill
		SET
			name = @name,
			level = @level
		WHERE skill_id = @id
	`

	result, err := tx.Exec(
		query,
		sql.Named("id", skill.SkillId),
		sql.Named("name", skill.Name),
		sql.Named("level", skill.Level),
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *ProfileCommandRepository) deleteSkill(tx *sql.Tx, id int) error {
	query := `
		DELETE skill
		WHERE skill_id = @id
	`
	result, err := tx.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) loadSocialMedia(tx *sql.Tx, profileId int) ([]domain.SocialMedia, error) {
	query := `
		  SELECT 
			 social_media_id
			,profile_id
			,platform
			,handle
			,link
  		FROM social_media
		WHERE profile_id = @profile_id
	`
	rows, err := tx.Query(query, sql.Named("profile_id", profileId))
	if err != nil {
		return []domain.SocialMedia{}, err
	}
	defer rows.Close()

	var socialMedias []domain.SocialMedia

	for rows.Next() {
		var socialMedia domain.SocialMedia

		err := rows.Scan(
			&socialMedia.SocialMediaId,
			&socialMedia.ProfileId,
			&socialMedia.Platform,
			&socialMedia.Handle,
			&socialMedia.Link,
		)
		if err != nil {
			return []domain.SocialMedia{}, err
		}

		socialMedias = append(socialMedias, socialMedia)
	}

	if err := rows.Err(); err != nil {
		return []domain.SocialMedia{}, err
	}

	return socialMedias, nil
}

func (p *ProfileCommandRepository) addSocialMedia(tx *sql.Tx, socialMedia *domain.SocialMedia) error {
	query := `
		INSERT INTO social_media (
			profile_id,
			platform,
			handle,
			link
		)
		OUTPUT INSERTED.social_media_id
		VALUES (
			@profile_id,
			@platform,
			@handle,
			@link
		)
	`

	var id int

	err := tx.QueryRow(
		query,
		sql.Named("profile_id", socialMedia.ProfileId),
		sql.Named("platform", socialMedia.Platform),
		sql.Named("handle", socialMedia.Handle),
		sql.Named("link", socialMedia.Link),
	).Scan(&id)

	if err != nil {
		return err
	}

	socialMedia.SocialMediaId = id

	return nil
}

func (p *ProfileCommandRepository) updateSocialMedia(tx *sql.Tx, socialMedia *domain.SocialMedia) error {
	query := `
		UPDATE social_media
		SET
			platform = @platform,
			handle = @handle,
			link = @link
		WHERE social_media_id = @id
	`

	result, err := tx.Exec(
		query,
		sql.Named("id", socialMedia.SocialMediaId),
		sql.Named("platform", socialMedia.Platform),
		sql.Named("handle", socialMedia.Handle),
		sql.Named("link", socialMedia.Link),
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *ProfileCommandRepository) deleteSocialMedia(tx *sql.Tx, id int) error {
	query := `
		DELETE social_media
		WHERE social_media_id = @id
	`
	result, err := tx.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *ProfileCommandRepository) saveProjects(tx *sql.Tx, profile *domain.Profile) error {
	if len(profile.Projects) > 0 {
		dbProjects, err := p.loadProjects(tx, profile.ProfileId)

		if err != nil {
			return err
		}

		dbMap := make(map[int]domain.Project)
		for _, p := range dbProjects {
			dbMap[p.ProjectId] = p
		}

		memMap := make(map[int]domain.Project)
		for _, p := range profile.Projects {
			if p.ProjectId != 0 {
				memMap[p.ProjectId] = p
			}
		}

		for id := range dbMap {
			if _, exists := memMap[id]; !exists {
				if err := p.deleteProject(tx, id); err != nil {
					return err
				}
			}
		}

		for _, project := range profile.Projects {
			if project.ProjectId == 0 {
				if err := p.addProject(tx, &project); err != nil {
					return err
				}
			} else {
				if err := p.updateProject(tx, &project); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (p *ProfileCommandRepository) saveEducations(tx *sql.Tx, profile *domain.Profile) error {
	if len(profile.Educations) > 0 {
		dbEducations, err := p.loadEducations(tx, profile.ProfileId)

		if err != nil {
			return err
		}

		dbMap := make(map[int]domain.Education)
		for _, education := range dbEducations {
			dbMap[education.EducationId] = education
		}

		memMap := make(map[int]domain.Education)
		for _, education := range profile.Educations {
			if education.EducationId != 0 {
				memMap[education.EducationId] = education
			}
		}

		for id := range dbMap {
			if _, exists := memMap[id]; !exists {
				if err := p.deleteEducation(tx, id); err != nil {
					return err
				}
			}
		}

		for _, education := range profile.Educations {
			if education.EducationId == 0 {
				if err := p.addEducation(tx, &education); err != nil {
					return err
				}
			} else {
				if err := p.updateEducation(tx, &education); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (p *ProfileCommandRepository) saveCertifications(tx *sql.Tx, profile *domain.Profile) error {
	if len(profile.Certifications) > 0 {
		dbCertifications, err := p.loadCertifications(tx, profile.ProfileId)

		if err != nil {
			return err
		}

		dbMap := make(map[int]domain.Certification)
		for _, certification := range dbCertifications {
			dbMap[certification.CertificationId] = certification
		}

		memMap := make(map[int]domain.Certification)
		for _, certification := range profile.Certifications {
			if certification.CertificationId != 0 {
				memMap[certification.CertificationId] = certification
			}
		}

		for id := range dbMap {
			if _, exists := memMap[id]; !exists {
				if err := p.deleteCertification(tx, id); err != nil {
					return err
				}
			}
		}

		for _, certification := range profile.Certifications {
			if certification.CertificationId == 0 {
				if err := p.addCertification(tx, &certification); err != nil {
					return err
				}
			} else {
				if err := p.updateCertification(tx, &certification); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (p *ProfileCommandRepository) saveExperiences(tx *sql.Tx, profile *domain.Profile) error {
	if len(profile.Certifications) > 0 {
		dbExperiences, err := p.loadExperiences(tx, profile.ProfileId)

		if err != nil {
			return err
		}

		dbMap := make(map[int]domain.Experience)
		for _, experience := range dbExperiences {
			dbMap[experience.ExperienceId] = experience
		}

		memMap := make(map[int]domain.Experience)
		for _, experience := range profile.Experiences {
			if experience.ExperienceId != 0 {
				memMap[experience.ExperienceId] = experience
			}
		}

		for id := range dbMap {
			if _, exists := memMap[id]; !exists {
				if err := p.deleteExperience(tx, id); err != nil {
					return err
				}
			}
		}

		for _, experience := range profile.Experiences {
			if experience.ExperienceId == 0 {
				if err := p.addExperience(tx, &experience); err != nil {
					return err
				}
			} else {
				if err := p.updateExperience(tx, &experience); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (p *ProfileCommandRepository) saveLanguages(tx *sql.Tx, profile *domain.Profile) error {
	if len(profile.Certifications) > 0 {
		dbLanguages, err := p.loadLanguages(tx, profile.ProfileId)

		if err != nil {
			return err
		}

		dbMap := make(map[int]domain.Language)
		for _, language := range dbLanguages {
			dbMap[language.LanguageId] = language
		}

		memMap := make(map[int]domain.Language)
		for _, language := range profile.Languages {
			if language.LanguageId != 0 {
				memMap[language.LanguageId] = language
			}
		}

		for id := range dbMap {
			if _, exists := memMap[id]; !exists {
				if err := p.deleteLanguage(tx, id); err != nil {
					return err
				}
			}
		}

		for _, language := range profile.Languages {
			if language.LanguageId == 0 {
				if err := p.addLanguage(tx, &language); err != nil {
					return err
				}
			} else {
				if err := p.updateLanguage(tx, &language); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (p *ProfileCommandRepository) saveSocialMedias(tx *sql.Tx, profile *domain.Profile) error {
	if len(profile.Certifications) > 0 {
		dbSocialMedias, err := p.loadSocialMedia(tx, profile.ProfileId)

		if err != nil {
			return err
		}

		dbMap := make(map[int]domain.SocialMedia)
		for _, socialMedia := range dbSocialMedias {
			dbMap[socialMedia.SocialMediaId] = socialMedia
		}

		memMap := make(map[int]domain.SocialMedia)
		for _, socialMedia := range profile.SocialMedias {
			if socialMedia.SocialMediaId != 0 {
				memMap[socialMedia.SocialMediaId] = socialMedia
			}
		}

		for id := range dbMap {
			if _, exists := memMap[id]; !exists {
				if err := p.deleteSocialMedia(tx, id); err != nil {
					return err
				}
			}
		}

		for _, socialMedia := range profile.SocialMedias {
			if socialMedia.SocialMediaId == 0 {
				if err := p.addSocialMedia(tx, &socialMedia); err != nil {
					return err
				}
			} else {
				if err := p.updateSocialMedia(tx, &socialMedia); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (p *ProfileCommandRepository) saveSkills(tx *sql.Tx, profile *domain.Profile) error {
	if len(profile.Certifications) > 0 {
		dbSkills, err := p.loadSkills(tx, profile.ProfileId)

		if err != nil {
			return err
		}

		dbMap := make(map[int]domain.Skill)
		for _, skill := range dbSkills {
			dbMap[skill.SkillId] = skill
		}

		memMap := make(map[int]domain.Skill)
		for _, skill := range profile.Skills {
			if skill.SkillId != 0 {
				memMap[skill.SkillId] = skill
			}
		}

		for id := range dbMap {
			if _, exists := memMap[id]; !exists {
				if err := p.deleteSkill(tx, id); err != nil {
					return err
				}
			}
		}

		for _, skill := range profile.Skills {
			if skill.SkillId == 0 {
				if err := p.addSkill(tx, &skill); err != nil {
					return err
				}
			} else {
				if err := p.updateSkill(tx, &skill); err != nil {
					return err
				}
			}
		}

	}

	return nil
}
