package dto

import "resume-genAI-api/cmd/api/model"

type ProfileDTO struct {
	ProfileId   int    `json:"profile_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Birthdate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`

	Projects       []model.Project       `json:"projects"`
	Certifications []model.Certification `json:"certifications"`
	Contacts       []model.Contact       `json:"contacts"`
	SocialMedias   []model.SocialMedia   `json:"socialMedias"`
	Educations     []model.Education     `json:"educations"`
	Experiences    []model.Experience    `json:"experiences"`
	Skills         []skillDTO            `json:"skills"`
	Languages      []model.Language      `json:"languages"`
}
