package model

type Certification struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Issuer     string `json:"issuer"`
	DateIssued string `json:"date_issued"`
}
