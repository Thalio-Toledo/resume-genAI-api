package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"resume-genAI-api/cmd/api/model"
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

func (p *ProfileRepository) LoadCertifications(profile *model.Profile) error {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
	}
	defer rows.Close()

	var certifications []model.Certification

	for rows.Next() {
		var certification model.Certification

		err := rows.Scan(
			&certification.Certification_Id,
			&certification.ProfileId,
			&certification.Name,
			&certification.Issuer,
			&certification.DateIssued,
		)
		if err != nil {
			return err
		}

		certifications = append(certifications, certification)

	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.Certifications = certifications
	return nil
}

func (p *ProfileRepository) LoadContacts(profile *model.Profile) error {
	query := `
		SELECT
			 contact_id
			,profile_id
			,email
			,phone_number
		FROM contact
		WHERE profile_id = @profile_id
	`
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
	}
	defer rows.Close()

	var contacts []model.Contact

	for rows.Next() {
		var contact model.Contact

		err := rows.Scan(
			&contact.ContactId,
			&contact.ProfileId,
			&contact.Email,
			&contact.PhoneNumber,
		)
		if err != nil {
			return err
		}

		contacts = append(contacts, contact)

	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.Contacts = contacts
	return nil
}

func (p *ProfileRepository) LoadEducations(profile *model.Profile) error {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
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
			return err
		}

		educations = append(educations, education)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.Educations = educations
	return nil
}

func (p *ProfileRepository) LoadExperiences(profile *model.Profile) error {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
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
			return err
		}

		experiences = append(experiences, experience)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.Experiences = experiences
	return nil
}

func (p *ProfileRepository) LoadLanguages(profile *model.Profile) error {
	query := `
		  SELECT 
			 language_id
			,profile_id
			,name
			,level
  		FROM language
		WHERE profile_id = @profile_id
	`
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
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
			return err
		}

		languages = append(languages, language)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.Languages = languages
	return nil
}

func (p *ProfileRepository) LoadSkill(profile *model.Profile) error {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
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
			return err
		}

		if skill.EmbeddingsJSON.Valid {
			err = json.Unmarshal([]byte(skill.EmbeddingsJSON.String), &skill.Embeddings)
			if err != nil {
				return err
			}
		}

		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.Skills = skills
	return nil
}

func (p *ProfileRepository) LoadSocialMedia(profile *model.Profile) error {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profile.ProfileId))
	if err != nil {
		return err
	}
	defer rows.Close()

	var socialMedias []model.SocialMedia

	for rows.Next() {
		var socialMedia model.SocialMedia

		err := rows.Scan(
			&socialMedia.SocialMediaId,
			&socialMedia.ProfileId,
			&socialMedia.Platform,
			&socialMedia.Handle,
			&socialMedia.Link,
		)
		if err != nil {
			return err
		}

		socialMedias = append(socialMedias, socialMedia)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	profile.SocialMedias = socialMedias
	return nil
}

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
