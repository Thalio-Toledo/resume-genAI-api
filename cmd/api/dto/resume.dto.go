package dto

import (
	"resume-genAI-api/cmd/api/domain"
)

type Resume struct {
	Profile        domain.Profile `json:"profile"`
	Match          float64        `json:"match"`
	SkillsRequired []domain.Skill `json:"skillsRequired"`
}
