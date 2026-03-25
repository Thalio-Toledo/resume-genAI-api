package dto

type skillDTO struct {
	SkillId   int    `json:"skill_id"`
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Level     string `json:"level"`
}
