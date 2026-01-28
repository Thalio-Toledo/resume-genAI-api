package model

type Contact struct {
	ContactId   int    `json:"contact_id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
