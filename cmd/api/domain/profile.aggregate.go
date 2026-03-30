package domain

import (
	"errors"
	ai "resume-genAI-api/cmd/api/ai"
)

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

func (p *Profile) UpdateProject(proj Project) error {
	for i, pro := range p.Projects {
		if pro.ProjectId == proj.ProjectId {
			p.Projects[i] = proj
			return nil
		}
	}
	return errors.New("Project not found for this profile")
}

func (p *Profile) UpdateCertification(cert Certification) error {
	for i, c := range p.Certifications {
		if c.CertificationId == cert.CertificationId {
			p.Certifications[i] = cert
			return nil
		}
	}
	return errors.New("Certification not found for this profile")
}

func (p *Profile) UpdateSocialMedia(soc SocialMedia) error {
	for i, s := range p.SocialMedias {
		if s.SocialMediaId == soc.SocialMediaId {
			p.SocialMedias[i] = soc
			return nil
		}
	}
	return errors.New("SocialMedia not found for this profile")
}

func (p *Profile) UpdateEducation(edu Education) error {
	for i, e := range p.Educations {
		if e.EducationId == edu.EducationId {
			p.Educations[i] = edu
			return nil
		}
	}
	return errors.New("Education not found for this profile")
}

func (p *Profile) UpdateExperience(ex Experience) error {
	for i, e := range p.Experiences {
		if e.ExperienceId == ex.ExperienceId {
			p.Experiences[i] = ex
			return nil
		}
	}
	return errors.New("Experience not found for this profile")
}

func (p *Profile) UpdateSkill(skill Skill) error {
	for i, sk := range p.Skills {
		if sk.SkillId == skill.SkillId {
			embeddings, _ := ai.GenerateEmbedding(skill.Name)
			skill.Embeddings = embeddings
			p.Skills[i] = skill
			return nil
		}
	}
	return errors.New("Skill not found for this profile")
}

func (p *Profile) UpdateLanguage(lang Language) error {
	for i, l := range p.Languages {
		if l.LanguageId == lang.LanguageId {
			p.Languages[i] = lang
			return nil
		}
	}
	return errors.New("Language not found for this profile")
}
