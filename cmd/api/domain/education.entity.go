package domain

type Education struct {
	EducationId int    `json:"education_id"`
	ProfileId   int    `json:"profile_id"`
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
