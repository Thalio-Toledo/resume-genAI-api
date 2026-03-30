package infrastructure

import "resume-genAI-api/cmd/api/domain"

type ProfileLoader struct {
	repo    *ProfileQueryRepository
	profile *domain.Profile
	err     error
}

func (l *ProfileLoader) LoadProjects() *ProfileLoader {
	if l.err != nil {
		return l
	}

	l.profile.Projects, l.err = l.repo.loadProjects(l.profile.ProfileId)
	return l
}

func (l *ProfileLoader) LoadCertifications() *ProfileLoader {
	if l.err != nil {
		return l
	}

	l.profile.Certifications, l.err = l.repo.loadCertifications(l.profile.ProfileId)
	return l
}

func (l *ProfileLoader) LoadExperiences() *ProfileLoader {
	if l.err != nil {
		return l
	}

	l.profile.Experiences, l.err = l.repo.loadExperiences(l.profile.ProfileId)
	return l
}

func (l *ProfileLoader) LoadEducations() *ProfileLoader {
	if l.err != nil {
		return l
	}

	l.profile.Educations, l.err = l.repo.loadEducations(l.profile.ProfileId)
	return l
}

func (l *ProfileLoader) LoadSocialMedias() *ProfileLoader {
	if l.err != nil {
		return l
	}

	l.profile.SocialMedias, l.err = l.repo.loadSocialMedias(l.profile.ProfileId)
	return l
}

func (l *ProfileLoader) LoadLanguages() *ProfileLoader {
	if l.err != nil {
		return l
	}

	l.profile.Languages, l.err = l.repo.loadLanguages(l.profile.ProfileId)
	return l
}

func (l *ProfileLoader) LoadSkills() *ProfileLoader {
	if l.err != nil {
		return l
	}

	l.profile.Skills, l.err = l.repo.loadSkills(l.profile.ProfileId)
	return l
}

func (l *ProfileLoader) Result() (*domain.Profile, error) {
	return l.profile, l.err
}
