package dto

import "resume-genAI-api/cmd/api/domain"

type ProfileDTO struct {
	ProfileId   int    `json:"profile_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Birthdate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`

	Projects       []domain.Project       `json:"projects"`
	Certifications []domain.Certification `json:"certifications"`
	Contacts       []domain.Contact       `json:"contacts"`
	SocialMedias   []domain.SocialMedia   `json:"socialMedias"`
	Educations     []domain.Education     `json:"educations"`
	Experiences    []domain.Experience    `json:"experiences"`
	Skills         []skillDTO             `json:"skills"`
	Languages      []domain.Language      `json:"languages"`
}
