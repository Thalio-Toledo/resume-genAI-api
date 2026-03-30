package application

import (
	"resume-genAI-api/cmd/api/domain"
	"resume-genAI-api/cmd/api/infrastructure"
)

type ProfileUseCase struct {
	repoCommand *infrastructure.ProfileCommandRepository
	repoQuery   *infrastructure.ProfileQueryRepository
}

func NewProfileUseCase(repoCommand *infrastructure.ProfileCommandRepository, repoQuery *infrastructure.ProfileQueryRepository) *ProfileUseCase {
	return &ProfileUseCase{repoCommand: repoCommand, repoQuery: repoQuery}
}

func (uc *ProfileUseCase) Get() ([]domain.Profile, error) {
	return uc.repoQuery.Get()
}

func (uc *ProfileUseCase) FindByID(id int) (*domain.Profile, error) {
	profile, _ := uc.repoQuery.
		FindProfile(id).
		LoadProjects().
		LoadCertifications().
		LoadEducations().
		LoadExperiences().
		LoadLanguages().
		LoadSkills().
		LoadSocialMedias().
		Result()

	return profile, nil
}

func (uc *ProfileUseCase) Create(profile *domain.Profile) error {
	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) Update(profile *domain.Profile) error {
	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) AddCertification(certification domain.Certification) error {
	profile, err := uc.FindByID(certification.ProfileId)
	if err != nil {
		return err
	}

	err = profile.AddCertification(certification)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) UpdateCertification(certification domain.Certification) error {
	profile, err := uc.FindByID(certification.ProfileId)
	if err != nil {
		return err
	}

	err = profile.UpdateCertification(certification)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) AddEducation(education domain.Education) error {
	profile, err := uc.FindByID(education.ProfileId)
	if err != nil {
		return err
	}

	err = profile.AddEducation(education)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) UpdateEducation(education domain.Education) error {
	profile, err := uc.FindByID(education.ProfileId)
	if err != nil {
		return err
	}

	err = profile.UpdateEducation(education)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) AddExperience(experience domain.Experience) error {
	profile, err := uc.FindByID(experience.ProfileId)
	if err != nil {
		return err
	}

	err = profile.AddExperience(experience)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) UpdateExperience(experience domain.Experience) error {
	profile, err := uc.FindByID(experience.ProfileId)
	if err != nil {
		return err
	}

	err = profile.UpdateExperience(experience)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) AddLanguage(language domain.Language) error {
	profile, err := uc.FindByID(language.ProfileId)
	if err != nil {
		return err
	}

	err = profile.AddLanguage(language)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) UpdateLanguage(language domain.Language) error {
	profile, err := uc.FindByID(language.ProfileId)
	if err != nil {
		return err
	}

	err = profile.UpdateLanguage(language)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) AddProject(project domain.Project) error {
	profile, err := uc.FindByID(project.ProfileId)
	if err != nil {
		return err
	}

	err = profile.AddProject(project)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) UpdateProject(project domain.Project) error {
	profile, err := uc.FindByID(project.ProfileId)
	if err != nil {
		return err
	}

	err = profile.UpdateProject(project)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) AddSkill(skill domain.Skill) error {
	profile, err := uc.FindByID(skill.ProfileId)
	if err != nil {
		return err
	}

	err = profile.AddSkill(skill)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) UpdateSkill(skill domain.Skill) error {
	profile, err := uc.FindByID(skill.ProfileId)
	if err != nil {
		return err
	}

	err = profile.UpdateSkill(skill)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) AddSocialMedia(socialMedia domain.SocialMedia) error {
	profile, err := uc.FindByID(socialMedia.ProfileId)
	if err != nil {
		return err
	}

	err = profile.AddSocialMedia(socialMedia)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) UpdateSocialMedia(socialMedia domain.SocialMedia) error {
	profile, err := uc.FindByID(socialMedia.ProfileId)
	if err != nil {
		return err
	}

	err = profile.UpdateSocialMedia(socialMedia)
	if err != nil {
		return err
	}

	return uc.repoCommand.Save(profile)
}

func (uc *ProfileUseCase) Delete(profileId int) error {
	return uc.repoCommand.Delete(profileId)
}
