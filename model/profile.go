package model

type ErrorResponse struct {
	Error string `json:"error"`
}
type Profile struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Birthdate   string `json:"birthdate"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`

	Projects       []Project       `json:"projects"`
	Certifications []Certification `json:"certifications"`
	Contacts       []Contact       `json:"contacts"`
	SocialMedias   []SocialMedia   `json:"social_medias"`
	Educations     []Education     `json:"educations"`
	Experiences    []Experience    `json:"experiences"`
	Skills         []Skill         `json:"skills"`
	Languages      []Language      `json:"languages"`
}
