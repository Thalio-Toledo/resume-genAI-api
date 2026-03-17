package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/cmd/api/model"
)

type ContactRepository struct {
	contacts []model.Contact
	db       *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{
		contacts: []model.Contact{},
		db:       db,
	}
}

func (r *ContactRepository) Get() ([]model.Contact, error) {
	query := `
		SELECT
			contact_id,
			profile_id,
			email,
			phone_number
		FROM contact
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.Contact

	for rows.Next() {
		var contact model.Contact

		err := rows.Scan(
			&contact.ContactId,
			&contact.ProfileId,
			&contact.Email,
			&contact.PhoneNumber,
		)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *ContactRepository) FindByID(id int) (*model.Contact, error) {
	query := `
		SELECT
			contact_id,
			profile_id,
			email,
			phone_number
		FROM contact
		WHERE contact_id = @id
	`
	var contact model.Contact

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&contact.ContactId,
		&contact.ProfileId,
		&contact.Email,
		&contact.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *ContactRepository) FindByProfileID(profileID int) ([]model.Contact, error) {
	query := `
		SELECT
			contact_id,
			profile_id,
			email,
			phone_number
		FROM contact
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.Contact

	for rows.Next() {
		var contact model.Contact

		err := rows.Scan(
			&contact.ContactId,
			&contact.ProfileId,
			&contact.Email,
			&contact.PhoneNumber,
		)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *ContactRepository) Create(contact model.Contact) (int, error) {
	query := `
		INSERT INTO contact (
			profile_id,
			email,
			phone_number
		)
		OUTPUT INSERTED.contact_id
		VALUES (
			@profile_id,
			@email,
			@phone_number
		)
	`

	var id int

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", contact.ProfileId),
		sql.Named("email", contact.Email),
		sql.Named("phone_number", contact.PhoneNumber),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ContactRepository) Update(contact model.Contact) (bool, error) {
	query := `
		UPDATE contact
		SET
			email = @email,
			phone_number = @phone_number
		WHERE contact_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", contact.ContactId),
		sql.Named("email", contact.Email),
		sql.Named("phone_number", contact.PhoneNumber),
	)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, sql.ErrNoRows
	}

	return true, nil
}

func (r *ContactRepository) Delete(id int) (bool, error) {
	query := `
		DELETE contact
		WHERE contact_id = @id
	`
	result, err := r.db.Exec(query, sql.Named("id", id))

	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, errors.New("no rows affected")
	}
	return true, nil
}
