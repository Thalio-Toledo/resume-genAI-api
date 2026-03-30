package domain

type Project struct {
	BaseModel
	ProjectId   int    `json:"project_id"`
	ProfileId   int    `json:"profile_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
