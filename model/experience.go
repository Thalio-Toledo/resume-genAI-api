package model

type Experience struct {
	ExperienceId string `json:"experience_id"`
	Company      string `json:"company"`
	IsCurrent    bool   `json:"is_current"`
	Role         string `json:"role"`
	Description  string `json:"description"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}
