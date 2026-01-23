package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type SocialMediaRepository struct {
	socialMedias []model.SocialMedia
}

func NewSocialMediaRepository() *SocialMediaRepository {
	return &SocialMediaRepository{
		socialMedias: []model.SocialMedia{},
	}
}

func (r *SocialMediaRepository) GetAll() []model.SocialMedia {
	return r.socialMedias
}

func (r *SocialMediaRepository) FindByHandle(handle string) (*model.SocialMedia, error) {
	for _, s := range r.socialMedias {
		if s.Handle == handle {
			return &s, nil
		}
	}
	return nil, errors.New("SocialMedia not found")
}

func (r *SocialMediaRepository) Create(sm model.SocialMedia) (string, error) {
	r.socialMedias = append(r.socialMedias, sm)
	return sm.Handle, nil
}

func (r *SocialMediaRepository) Update(sm model.SocialMedia) (bool, error) {
	for i, s := range r.socialMedias {
		if s.Handle == sm.Handle {
			r.socialMedias[i] = sm
			return true, nil
		}
	}
	return false, errors.New("SocialMedia not found")
}

func (r *SocialMediaRepository) Delete(handle string) (bool, error) {
	for i, s := range r.socialMedias {
		if s.Handle == handle {
			r.socialMedias = append(r.socialMedias[:i], r.socialMedias[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("SocialMedia not found")
}
