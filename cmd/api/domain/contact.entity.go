package domain

type Contact struct {
	ContactId   int    `json:"contact_id"`
	ProfileId   int    `json:"profile_id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
