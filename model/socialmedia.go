package model

type SocialMedia struct {
	SocialMediaId int    `json:"social_media_id"`
	ProfileId     int    `json:"profile_id"`
	Platform      string `json:"platform"`
	Handle        string `json:"handle"`
	Link          string `json:"link"`
}
