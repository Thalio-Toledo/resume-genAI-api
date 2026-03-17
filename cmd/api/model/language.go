package model

type Language struct {
	LanguageId string `json:"language_id"`
	ProfileID  int    `json:"profile_id"`
	Name       string `json:"name"`
	Level      string `json:"level"`
}
