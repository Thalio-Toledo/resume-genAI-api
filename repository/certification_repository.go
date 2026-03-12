package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/model"
)

type CertificationRepository struct {
	certifications []model.Certification
	db             *sql.DB
}

func NewCertificationRepository(db *sql.DB) *CertificationRepository {
	return &CertificationRepository{
		certifications: []model.Certification{},
		db:             db,
	}
}

func (r *CertificationRepository) Get() ([]model.Certification, error) {
	query := `
		SELECT
			 certification_id
			,profile_id
			,name
			,issuer
			,date_issued
		FROM certification
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var certifications []model.Certification

	for rows.Next() {
		var certification model.Certification

		err := rows.Scan(
			&certification.Certification_Id,
			&certification.ProfileId,
			&certification.Name,
			&certification.Issuer,
			&certification.DateIssued,
		)
		if err != nil {
			return nil, err
		}

		certifications = append(certifications, certification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return certifications, nil
}

func (r *CertificationRepository) FindByID(id int) (*model.Certification, error) {
	query := `
		SELECT
			 certification_id
			,profile_id
			,name
			,issuer
			,date_issued
		FROM certification
		WHERE certification_id = @id
	`
	var certification model.Certification

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&certification.Certification_Id,
		&certification.ProfileId,
		&certification.Name,
		&certification.Issuer,
		&certification.DateIssued,
	)
	if err != nil {
		return nil, err
	}

	return &certification, nil
}

func (r *CertificationRepository) FindByProfileID(profileID int) ([]model.Certification, error) {
	query := `
		SELECT
			 certification_id
			,profile_id
			,name
			,issuer
			,date_issued
		FROM certification
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var certifications []model.Certification

	for rows.Next() {
		var certification model.Certification

		err := rows.Scan(
			&certification.Certification_Id,
			&certification.ProfileId,
			&certification.Name,
			&certification.Issuer,
			&certification.DateIssued,
		)
		if err != nil {
			return nil, err
		}

		certifications = append(certifications, certification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return certifications, nil
}

func (r *CertificationRepository) Create(certification model.Certification) (int, error) {
	query := `
		INSERT INTO certification (
			profile_id,
			name,
			issuer,
			date_issued
		)
		OUTPUT INSERTED.certification_id
		VALUES (
			@profile_id,
			@name,
			@issuer,
			@date_issued
		)
	`

	var id int

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", certification.ProfileId),
		sql.Named("name", certification.Name),
		sql.Named("issuer", certification.Issuer),
		sql.Named("date_issued", certification.DateIssued),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CertificationRepository) Update(certification model.Certification) (bool, error) {
	query := `
		UPDATE certification
		SET
			name = @name,
			issuer = @issuer,
			date_issued = @date_issued
		WHERE certification_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", certification.Certification_Id),
		sql.Named("name", certification.Name),
		sql.Named("issuer", certification.Issuer),
		sql.Named("date_issued", certification.DateIssued),
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

func (r *CertificationRepository) Delete(id int) (bool, error) {
	query := `
		DELETE certification
		WHERE certification_id = @id
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
