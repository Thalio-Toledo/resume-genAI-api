package domain

import "database/sql"

type Skill struct {
	SkillId        int            `json:"skill_id"`
	ProfileId      int            `json:"profile_id"`
	Name           string         `json:"name"`
	Level          string         `json:"level"`
	EmbeddingsJSON sql.NullString `json:"embeddings" swaggerignore:"true"`
	Embeddings     []float32
}
