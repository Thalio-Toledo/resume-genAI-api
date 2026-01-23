package repository

import (
	"errors"
	"resume-genAI-api/model"
)

type ContactRepository struct {
	contacts []model.Contact
}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{
		contacts: []model.Contact{},
	}
}

func (r *ContactRepository) GetAll() []model.Contact {
	return r.contacts
}

func (r *ContactRepository) FindByEmail(email string) (*model.Contact, error) {
	for _, c := range r.contacts {
		if c.Email == email {
			return &c, nil
		}
	}
	return nil, errors.New("Contact not found")
}

func (r *ContactRepository) Create(contact model.Contact) (string, error) {
	r.contacts = append(r.contacts, contact)
	return contact.Email, nil
}

func (r *ContactRepository) Update(contact model.Contact) (bool, error) {
	for i, c := range r.contacts {
		if c.Email == contact.Email {
			r.contacts[i] = contact
			return true, nil
		}
	}
	return false, errors.New("Contact not found")
}

func (r *ContactRepository) Delete(email string) (bool, error) {
	for i, c := range r.contacts {
		if c.Email == email {
			r.contacts = append(r.contacts[:i], r.contacts[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Contact not found")
}
