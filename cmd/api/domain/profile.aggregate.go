package domain

import (
	"errors"
	ai "resume-genAI-api/cmd/api/AI"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
type Profile struct {
	BaseModel
	ProfileId   int    `json:"profile_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Birthdate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`

	Projects       []Project       `json:"projects"`
	Certifications []Certification `json:"certifications"`
	SocialMedias   []SocialMedia   `json:"socialMedias"`
	Educations     []Education     `json:"educations"`
	Experiences    []Experience    `json:"experiences"`
	Skills         []Skill         `json:"skills"`
	Languages      []Language      `json:"languages"`
}

func (p *Profile) AddProject(proj Project) error {
	for _, p := range p.Projects {
		if p.Name == proj.Name {
			return errors.New("Project already exists for this profile")
		}
	}

	p.Projects = append(p.Projects, proj)
	return nil
}

func (p *Profile) AddCertification(cert Certification) error {
	for _, c := range p.Certifications {
		if c.Name == cert.Name {
			return errors.New("Certification already exists for this profile")
		}
	}

	p.Certifications = append(p.Certifications, cert)
	return nil
}

func (p *Profile) AddSocialMedia(soc SocialMedia) error {
	p.SocialMedias = append(p.SocialMedias, soc)
	return nil
}

func (p *Profile) AddEducation(edu Education) error {
	p.Educations = append(p.Educations, edu)
	return nil
}

func (p *Profile) AddExperience(ex Experience) error {
	for _, exe := range p.Experiences {
		if exe.Company == ex.Company {
			return errors.New("Experience already exists for this profile")
		}
	}
	p.Experiences = append(p.Experiences, ex)
	return nil
}

func (p *Profile) AddSkill(skill Skill) error {
	for _, sk := range p.Skills {
		if sk.Name == skill.Name {
			return errors.New("Skill already exists for this profile")
		}
	}
	embeddings, _ := ai.GenerateEmbedding(skill.Name)
	skill.Embeddings = embeddings
	p.Skills = append(p.Skills, skill)
	return nil
}

func (p *Profile) AddLanguage(lang Language) error {
	for _, l := range p.Languages {
		if l.Name == lang.Name {
			return errors.New("Language already exists for this profile")
		}
	}
	p.Languages = append(p.Languages, lang)
	return nil
}
