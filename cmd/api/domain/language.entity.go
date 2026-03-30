package domain

type Language struct {
	LanguageId int    `json:"language_id"`
	ProfileId  int    `json:"profile_id"`
	Name       string `json:"name"`
	Level      string `json:"level"`
}
