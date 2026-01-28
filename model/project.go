package model

type Project struct {
	BaseModel
	ProjectId   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
