package useCase

import (
	"fmt"
	ai "resume-genAI-api/cmd/api/AI"
	skillMatch "resume-genAI-api/cmd/api/utils"
	"resume-genAI-api/model"
	"resume-genAI-api/repository"
)

type ProfileUseCase struct {
	repo *repository.ProfileRepository
}

func NewProfileUseCase(repo *repository.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{repo: repo}
}

func (uc *ProfileUseCase) Get() ([]model.Profile, error) {
	return uc.repo.Get()
}

func (uc *ProfileUseCase) FindByID(id int) (*model.Profile, error) {
	profile, _ := uc.repo.FindByID(id)
	uc.repo.LoadEducations(profile)
	uc.repo.LoadProjects(profile)
	uc.repo.LoadCertifications(profile)
	uc.repo.LoadContacts(profile)
	uc.repo.LoadExperiences(profile)
	uc.repo.LoadSkill(profile)
	uc.repo.LoadLanguages(profile)
	uc.repo.LoadSocialMedia(profile)

	vagaEmbedding, _ := ai.GenerateEmbedding("angularjs")

	var skills []model.Skill

	for _, skill := range profile.Skills {
		score := skillMatch.CosineSimilarity(vagaEmbedding, skill.Embeddings)

		if score >= 0.75 {
			skills = append(skills, skill)
			fmt.Printf("MATCH: %s (%.2f)\n", skill.Name, score)
		}
	}

	profile.Skills = skills
	return profile, nil
}

func (uc *ProfileUseCase) Create(profile model.Profile) (int, error) {
	return uc.repo.Create(profile)
}

func (uc *ProfileUseCase) Update(profile model.Profile) (bool, error) {
	return uc.repo.Update(profile)
}

func (uc *ProfileUseCase) Delete(id int) (bool, error) {
	return uc.repo.Delete(id)
}
