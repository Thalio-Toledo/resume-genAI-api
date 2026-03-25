package domain

type Education struct {
	EducationId string `json:"education_id"`
	ProfileID   int    `json:"profile_id"`
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
