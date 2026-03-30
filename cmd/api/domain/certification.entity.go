package domain

type Certification struct {
	CertificationId int    `json:"certification_id"`
	ProfileId       int    `json:"profile_id"`
	Name            string `json:"name"`
	Issuer          string `json:"issuer"`
	DateIssued      string `json:"date_issued"`

	Skills []Skill `json:"skills"`
}
