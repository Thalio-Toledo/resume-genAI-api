package model

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
	Contacts       []Contact       `json:"contacts"`
	SocialMedias   []SocialMedia   `json:"social_medias"`
	Educations     []Education     `json:"educations"`
	Experiences    []Experience    `json:"experiences"`
	Skills         []Skill         `json:"skills"`
	Languages      []Language      `json:"languages"`
}
