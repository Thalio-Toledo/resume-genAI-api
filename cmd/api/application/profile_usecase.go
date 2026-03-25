package application

import (
	"fmt"
	ai "resume-genAI-api/cmd/api/AI"
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/dto"
	"resume-genAI-api/cmd/api/infrastructure"
	skillMatch "resume-genAI-api/cmd/api/utils"
)

type ProfileUseCase struct {
	repo *infrastructure.ProfileRepository
}

func NewProfileUseCase(repo *infrastructure.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{repo: repo}
}

func (uc *ProfileUseCase) Get() ([]domain.Profile, error) {
	return uc.repo.Get()
}

func (uc *ProfileUseCase) FindByID(id int) (*domain.Profile, error) {
	profile, _ := uc.repo.FindByID(id)
	uc.repo.LoadEducations(profile)
	uc.repo.LoadProjects(profile)
	uc.repo.LoadCertifications(profile)
	uc.repo.LoadContacts(profile)
	uc.repo.LoadExperiences(profile)
	uc.repo.LoadSkill(profile)
	uc.repo.LoadLanguages(profile)
	uc.repo.LoadSocialMedia(profile)

	// vagaEmbedding, _ := ai.GenerateEmbedding("angularjs")

	// var skills []model.Skill

	// for _, skill := range profile.Skills {
	// 	score := skillMatch.CosineSimilarity(vagaEmbedding, skill.Embeddings)

	// 	if score >= 0.75 {
	// 		skills = append(skills, skill)
	// 		fmt.Printf("MATCH: %s (%.2f)\n", skill.Name, score)
	// 	}
	// }

	// profile.Skills = skills
	return profile, nil
}

func (uc *ProfileUseCase) Generate(job_description dto.RoleDescription) (*dto.Resume, error) {
	profile, err := uc.FindByID(job_description.ProfileId)
	skillsStrings, err := ai.Generate(job_description.Description)
	if err != nil {
		return nil, err
	}

	var skills []domain.Skill
	var skillsRequired []domain.Skill

	for _, skillString := range skillsStrings {
		var skill domain.Skill
		embeddings, err := ai.GenerateEmbedding(skillString)
		if err != nil {
			return nil, err
		}
		skill.Name = skillString
		skill.Embeddings = embeddings

		for i, skillProfile := range profile.Skills {
			score := skillMatch.CosineSimilarity(skill.Embeddings, skillProfile.Embeddings)
			if score >= 0.75 {
				skills = append(skills, skillProfile)
				fmt.Printf("MATCH: %s (%.2f)\n", skill.Name, score)
				break
			}

			if i == len(profile.Skills)-1 {
				skillsRequired = append(skillsRequired, skill)
			}
		}
	}

	profile.Skills = skills

	var resume dto.Resume
	resume.Profile = *profile
	resume.SkillsRequired = skillsRequired
	sk := len(profile.Skills)
	skr := len(skillsRequired)
	resume.Match = (float64(sk) / float64(skr)) * 100

	return &resume, nil
}

func (uc *ProfileUseCase) Create(profile domain.Profile) (int, error) {
	return uc.repo.Create(profile)
}

func (uc *ProfileUseCase) Update(profile domain.Profile) (bool, error) {
	return uc.repo.Update(profile)
}

func (uc *ProfileUseCase) Delete(id int) (bool, error) {
	return uc.repo.Delete(id)
}
