package repository

import (
	"database/sql"
	"errors"
	"resume-genAI-api/model"
)

type SocialMediaRepository struct {
	socialMedias []model.SocialMedia
	db           *sql.DB
}

func NewSocialMediaRepository(db *sql.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		socialMedias: []model.SocialMedia{},
		db:           db,
	}
}

func (r *SocialMediaRepository) Get() ([]model.SocialMedia, error) {
	query := `
		SELECT
			social_media_id,
			profile_id,
			platform,
			handle,
			link
		FROM social_media
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var socialMedias []model.SocialMedia

	for rows.Next() {
		var socialMedia model.SocialMedia

		err := rows.Scan(
			&socialMedia.SocialMediaId,
			&socialMedia.ProfileId,
			&socialMedia.Platform,
			&socialMedia.Handle,
			&socialMedia.Link,
		)
		if err != nil {
			return nil, err
		}

		socialMedias = append(socialMedias, socialMedia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (r *SocialMediaRepository) FindByID(id int) (*model.SocialMedia, error) {
	query := `
		SELECT
			social_media_id,
			profile_id,
			platform,
			handle,
			link
		FROM social_media
		WHERE social_media_id = @id
	`
	var socialMedia model.SocialMedia

	err := r.db.QueryRow(query, sql.Named("id", id)).Scan(
		&socialMedia.SocialMediaId,
		&socialMedia.ProfileId,
		&socialMedia.Platform,
		&socialMedia.Handle,
		&socialMedia.Link,
	)
	if err != nil {
		return nil, err
	}

	return &socialMedia, nil
}

func (r *SocialMediaRepository) FindByProfileID(profileID int) ([]model.SocialMedia, error) {
	query := `
		SELECT
			social_media_id,
			profile_id,
			platform,
			handle,
			link
		FROM social_media
		WHERE profile_id = @profile_id
	`
	rows, err := r.db.Query(query, sql.Named("profile_id", profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var socialMedias []model.SocialMedia

	for rows.Next() {
		var socialMedia model.SocialMedia

		err := rows.Scan(
			&socialMedia.SocialMediaId,
			&socialMedia.ProfileId,
			&socialMedia.Platform,
			&socialMedia.Handle,
			&socialMedia.Link,
		)
		if err != nil {
			return nil, err
		}

		socialMedias = append(socialMedias, socialMedia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (r *SocialMediaRepository) Create(socialMedia model.SocialMedia) (int, error) {
	query := `
		INSERT INTO social_media (
			profile_id,
			platform,
			handle,
			link
		)
		OUTPUT INSERTED.social_media_id
		VALUES (
			@profile_id,
			@platform,
			@handle,
			@link
		)
	`

	var id int

	err := r.db.QueryRow(
		query,
		sql.Named("profile_id", socialMedia.ProfileId),
		sql.Named("platform", socialMedia.Platform),
		sql.Named("handle", socialMedia.Handle),
		sql.Named("link", socialMedia.Link),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SocialMediaRepository) Update(socialMedia model.SocialMedia) (bool, error) {
	query := `
		UPDATE social_media
		SET
			platform = @platform,
			handle = @handle,
			link = @link
		WHERE social_media_id = @id
	`

	result, err := r.db.Exec(
		query,
		sql.Named("id", socialMedia.SocialMediaId),
		sql.Named("platform", socialMedia.Platform),
		sql.Named("handle", socialMedia.Handle),
		sql.Named("link", socialMedia.Link),
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

func (r *SocialMediaRepository) Delete(id int) (bool, error) {
	query := `
		DELETE social_media
		WHERE social_media_id = @id
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
