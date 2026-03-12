package model

type Skill struct {
	SkillId   string `json:"skill_id"`
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Level     string `json:"level"`
}
