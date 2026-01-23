package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type LanguageRepository struct {
	languages []model.Language
}

func NewLanguageRepository() *LanguageRepository {
	return &LanguageRepository{
		languages: []model.Language{},
	}
}

func (r *LanguageRepository) GetAll() []model.Language {
	return r.languages
}

func (r *LanguageRepository) FindByID(id string) (*model.Language, error) {
	for _, l := range r.languages {
		if l.ID == id {
			return &l, nil
		}
	}
	return nil, errors.New("Language not found")
}

func (r *LanguageRepository) Create(lang model.Language) (string, error) {
	r.languages = append(r.languages, lang)
	return lang.ID, nil
}

func (r *LanguageRepository) Update(lang model.Language) (bool, error) {
	for i, l := range r.languages {
		if l.ID == lang.ID {
			r.languages[i] = lang
			return true, nil
		}
	}
	return false, errors.New("Language not found")
}

func (r *LanguageRepository) Delete(id string) (bool, error) {
	for i, l := range r.languages {
		if l.ID == id {
			r.languages = append(r.languages[:i], r.languages[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Language not found")
}
