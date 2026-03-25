package application

import (
	"fmt"
	ai "resume-genAI-api/cmd/api/AI"
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/dto"
	"resume-genAI-api/cmd/api/infrastructure"
	skillMatch "resume-genAI-api/cmd/api/utils"
)

type GenerateResumeUseCase struct {
	repo               *infrastructure.ProfileRepository
	findProfileUseCase *FindProfileUseCase
}

func NewGenerateResumeUseCase(repo *infrastructure.ProfileRepository, findUC *FindProfileUseCase) *GenerateResumeUseCase {
	return &GenerateResumeUseCase{repo: repo, findProfileUseCase: findUC}
}

func (uc *GenerateResumeUseCase) Generate(job_description dto.RoleDescription) (*dto.Resume, error) {
	profile, err := uc.findProfileUseCase.FindByID(job_description.ProfileId)
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
