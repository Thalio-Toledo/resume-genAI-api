package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type CertificationRepository struct {
	certifications []model.Certification
}

func NewCertificationRepository() *CertificationRepository {
	return &CertificationRepository{
		certifications: []model.Certification{},
	}
}

func (r *CertificationRepository) GetAll() ([]model.Certification, error) {
	return r.certifications, nil
}

func (r *CertificationRepository) FindByID(id string) (*model.Certification, error) {
	for _, c := range r.certifications {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("Certification not found")
}

func (r *CertificationRepository) Create(cert model.Certification) (string, error) {
	r.certifications = append(r.certifications, cert)
	return cert.ID, nil
}

func (r *CertificationRepository) Update(cert model.Certification) (bool, error) {
	for i, c := range r.certifications {
		if c.ID == cert.ID {
			r.certifications[i] = cert
			return true, nil
		}
	}
	return false, errors.New("Certification not found")
}

func (r *CertificationRepository) Delete(id string) (bool, error) {
	for i, c := range r.certifications {
		if c.ID == id {
			r.certifications = append(r.certifications[:i], r.certifications[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Certification not found")
}
