package model

type Certification struct {
	Certification_Id int    `json:"certification_id"`
	Name             string `json:"name"`
	Issuer           string `json:"issuer"`
	DateIssued       string `json:"date_issued"`
}
