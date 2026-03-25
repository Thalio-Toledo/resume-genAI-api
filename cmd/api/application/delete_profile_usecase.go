package application

import "resume-genAI-api/cmd/api/infrastructure"

type DeleteProfileUseCase struct {
	repo *infrastructure.ProfileRepository
}

func NewDeleteProfileUseCase(repo *infrastructure.ProfileRepository) *DeleteProfileUseCase {
	return &DeleteProfileUseCase{repo: repo}
}

func (uc *ProfileUseCase) Delete(id int) (bool, error) {
	return uc.repo.Delete(id)
}
