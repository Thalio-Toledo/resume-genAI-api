package useCase

import (
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type CertificationUseCase struct {
	repo *repository.CertificationRepository
}

func NewCertificationUseCase(repo *repository.CertificationRepository) *CertificationUseCase {
	return &CertificationUseCase{repo: repo}
}

func (uc *CertificationUseCase) GetAll() ([]model.Certification, error) {
	return uc.repo.GetAll()
}

func (uc *CertificationUseCase) FindByID(id int) (*model.Certification, error) {
	return uc.repo.FindByID(id)
}

func (uc *CertificationUseCase) Create(cert model.Certification) (int, error) {
	return uc.repo.Create(cert)
}

func (uc *CertificationUseCase) Update(cert model.Certification) (bool, error) {
	return uc.repo.Update(cert)
}

func (uc *CertificationUseCase) Delete(id int) (bool, error) {
	return uc.repo.Delete(id)
}
