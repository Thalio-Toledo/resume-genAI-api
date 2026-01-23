package model

type Resume struct {
	ID        string `json:"id"`
	ProfileID string `json:"profile_id"`

	SkillIDs         []string `json:"skill_ids"`         // IDs das skills selecionadas
	ExperienceIDs    []string `json:"experience_ids"`    // IDs das experiências selecionadas
	ProjectIDs       []string `json:"project_ids"`       // IDs dos projetos selecionados
	CertificationIDs []string `json:"certification_ids"` // IDs das certificações selecionadas
	EducationIDs     []string `json:"education_ids"`     // IDs das formações selecionadas
	LanguageIDs      []string `json:"language_ids"`      // IDs dos idiomas selecionados
	SocialMediaIDs   []string `json:"social_media_ids"`  // IDs das redes sociais selecionadas
	ContactIDs       []string `json:"contact_ids"`       // IDs dos contatos selecionados
}
