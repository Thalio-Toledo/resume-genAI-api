package model

// SkillDTO is a DTO for Swagger documentation purposes
// Use this structure for API responses to avoid sql.NullString issues in Swagger generation
type SkillDTO struct {
	SkillId    string    `json:"skill_id"`
	ProfileID  int       `json:"profile_id"`
	Name       string    `json:"name"`
	Level      string    `json:"level"`
	Embeddings []float32 `json:"embeddings"`
}

// ProfileDTO is a DTO for Swagger documentation purposes
type ProfileDTO struct {
	ProfileId      int             `json:"profile_id"`
	Name           string          `json:"name"`
	Email          string          `json:"email"`
	Birthdate      string          `json:"birth_date"`
	PhoneNumber    string          `json:"phone_number"`
	Description    string          `json:"description"`
	Active         bool            `json:"active"`
	CreatedAt      string          `json:"created_at"`
	UpdatedAt      string          `json:"updated_at"`
	DeletedAt      string          `json:"deleted_at,omitempty"`
	Projects       []Project       `json:"projects"`
	Certifications []Certification `json:"certifications"`
	Contacts       []Contact       `json:"contacts"`
	SocialMedias   []SocialMedia   `json:"socialMedias"`
	Educations     []Education     `json:"educations"`
	Experiences    []Experience    `json:"experiences"`
	Skills         []SkillDTO      `json:"skills"`
	Languages      []Language      `json:"languages"`
}
