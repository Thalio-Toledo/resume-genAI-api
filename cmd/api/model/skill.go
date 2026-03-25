package model

import "database/sql"

type Skill struct {
	SkillId        int            `json:"skill_id"`
	ProfileID      int            `json:"profile_id"`
	Name           string         `json:"name"`
	Level          string         `json:"level"`
	EmbeddingsJSON sql.NullString `json:"embeddings"`
	Embeddings     []float32
}
