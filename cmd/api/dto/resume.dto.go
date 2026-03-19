package dto

import (
	"resume-genAI-api/cmd/api/model"
)

type Resume struct {
	Profile        model.Profile `json:"profile"`
	Match          float64       `json:"match"`
	SkillsRequired []model.Skill `json:"skillsRequired"`
}
