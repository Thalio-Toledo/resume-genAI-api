package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type ProfileRepository struct {
	profiles []model.Profile
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{
		profiles: []model.Profile{},
	}
}

func (p *ProfileRepository) Get() ([]model.Profile, error) {
	return p.profiles, nil
}

func (p *ProfileRepository) FindByID(id int) (*model.Profile, error) {
	for _, profile := range p.profiles {
		if profile.ID == id {
			return &profile, nil
		}
	}
	return nil, nil
}

func (p *ProfileRepository) Create(profile model.Profile) (int, error) {
	id := 1

	for i, _ := range p.profiles {
		id = i + 1
	}

	profile.ID = id + 1

	p.profiles = append(p.profiles, profile)
	return profile.ID, nil
}

func (p *ProfileRepository) Update(profile model.Profile) (bool, error) {
	for i, existingProfile := range p.profiles {
		if existingProfile.ID == profile.ID {
			p.profiles[i] = profile
			return true, nil
		}
	}
	return false, errors.New("Profile not found")
}

func (p *ProfileRepository) Delete(id int) (bool, error) {
	for i, profile := range p.profiles {
		if profile.ID == id {
			p.profiles = append(p.profiles[:i], p.profiles[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Profile not found")
}
