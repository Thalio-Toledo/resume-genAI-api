package infrastructure

import (
	"database/sql"
	"encoding/json"
	"resume-genAI-api/cmd/api/domain"
)

type ProfileQueryRepository struct {
	db *sql.DB
}

func NewProfileQueryRepository(db *sql.DB) *ProfileQueryRepository {
	return &ProfileQueryRepository{db: db}
}

func (p *ProfileQueryRepository) Get() ([]domain.Profile, error) {
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

	var profiles []domain.Profile

	for rows.Next() {
		var profile domain.Profile

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

func (p *ProfileQueryRepository) findByID(id int) (*domain.Profile, error) {
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
	var profile domain.Profile

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

func (r *ProfileQueryRepository) FindProfile(id int) *ProfileLoader {
	profile, err := r.findByID(id)

	return &ProfileLoader{
		repo:    r,
		profile: profile,
		err:     err,
	}
}

func (p *ProfileQueryRepository) loadProjects(profileId int) ([]domain.Project, error) {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profileId))
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

func (p *ProfileQueryRepository) loadCertifications(profileId int) ([]domain.Certification, error) {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profileId))
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

func (p *ProfileQueryRepository) loadEducations(profileId int) ([]domain.Education, error) {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profileId))
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

func (p *ProfileQueryRepository) loadExperiences(profileId int) ([]domain.Experience, error) {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profileId))
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

func (p *ProfileQueryRepository) loadLanguages(profileId int) ([]domain.Language, error) {
	query := `
		  SELECT 
			 language_id
			,profile_id
			,name
			,level
  		FROM language
		WHERE profile_id = @profile_id
	`
	rows, err := p.db.Query(query, sql.Named("profile_id", profileId))
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

func (p *ProfileQueryRepository) loadSkills(profileId int) ([]domain.Skill, error) {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profileId))
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

func (p *ProfileQueryRepository) loadSocialMedias(profileId int) ([]domain.SocialMedia, error) {
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
	rows, err := p.db.Query(query, sql.Named("profile_id", profileId))
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
