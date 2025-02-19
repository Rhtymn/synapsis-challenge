package repository

import (
	"context"
	"database/sql"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/jackc/pgx/v5"
)

type userRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepositoryPostgres {
	return &userRepositoryPostgres{
		db: db,
	}
}

func (r *userRepositoryPostgres) GetById(ctx context.Context, id int64) (domain.User, error) {
	user := domain.User{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"id": id,
	}
	query := `
		SELECT ` + constants.UserColumns + `
		FROM users 
		WHERE id = @id
			AND deleted_at IS NULL
	`

	var dateOfBirth sql.NullTime
	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&user.ID,
			&user.Name,
			&user.PhotoURL,
			&dateOfBirth,
			&user.Gender,
			&user.Account.ID,
			&user.MainAddressID,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, apperror.NewNotFound(err, "user not found")
		}
		return user, apperror.Wrap(err)
	}
	user.DateOfBirth = toTimePtr(dateOfBirth)

	return user, nil
}

func (r *userRepositoryPostgres) GetByIdAndLock(ctx context.Context, id int64) (domain.User, error) {
	user := domain.User{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"id": id,
	}
	query := `
		SELECT ` + constants.UserColumns + `
		FROM users 
		WHERE id = @id 
			AND deleted_at IS NULL 
		FOR UPDATE
	`

	var dateOfBirth sql.NullTime
	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&user.ID,
			&user.Name,
			&user.PhotoURL,
			&dateOfBirth,
			&user.Gender,
			&user.Account.ID,
			&user.MainAddressID,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, apperror.NewNotFound(err, "user not found")
		}
		return user, apperror.Wrap(err)
	}
	user.DateOfBirth = toTimePtr(dateOfBirth)

	return user, nil
}

func (r *userRepositoryPostgres) GetByAccountID(ctx context.Context, accountID int64) (domain.User, error) {
	user := domain.User{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"accountID": accountID,
	}
	query := `
		SELECT ` + constants.UserColumns + `
		FROM users 
		WHERE id_account = @accountID
			AND deleted_at IS NULL
	`

	var dateOfBirth sql.NullTime
	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&user.ID,
			&user.Name,
			&user.PhotoURL,
			&dateOfBirth,
			&user.Gender,
			&user.PhoneNumber,
			&user.Account.ID,
			&user.MainAddressID,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, apperror.NewNotFound(err, "user not found")
		}
		return user, apperror.Wrap(err)
	}
	user.DateOfBirth = toTimePtr(dateOfBirth)

	return user, nil
}

func (r *userRepositoryPostgres) Add(ctx context.Context, user domain.User) (domain.User, error) {
	u := domain.User{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"name":        user.Name,
		"photo":       user.PhotoURL,
		"dob":         user.DateOfBirth,
		"gender":      user.Gender,
		"phoneNumber": user.PhoneNumber,
		"accountId":   user.Account.ID,
	}
	query := `
		INSERT INTO users(name, photo_url, date_of_birth, gender, phone_number, id_account) 
			VALUES (@name, @photo, @dob, @gender, @phoneNumber, @accountId) 
		RETURNING ` + constants.UserColumns + `
	`

	var dateOfBirth sql.NullTime
	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&u.ID,
			&u.Name,
			&u.PhotoURL,
			&dateOfBirth,
			&u.Gender,
			&u.PhoneNumber,
			&u.Account.ID,
			&u.MainAddressID,
		)
	if err != nil {
		return u, apperror.Wrap(err)
	}
	user.DateOfBirth = toTimePtr(dateOfBirth)

	return u, nil
}

func (r *userRepositoryPostgres) Update(ctx context.Context, user domain.User) (domain.User, error) {
	u := domain.User{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"name":        user.Name,
		"photo":       user.PhotoURL,
		"dob":         user.DateOfBirth,
		"gender":      user.Gender,
		"phoneNumber": user.PhoneNumber,
		"id":          user.ID,
	}
	query := `
		UPDATE users 
			SET name = @name, 
				photo_url = @photo, 
				date_of_birth = @dob, 
				gender = @gender, 
				phone_number = @phoneNumber
			WHERE id = @id 
		RETURNING ` + constants.UserColumns + `
	`

	var dateOfBirth sql.NullTime
	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&u.ID,
			&u.Name,
			&u.PhotoURL,
			&dateOfBirth,
			&u.Gender,
			&u.PhoneNumber,
			&u.Account.ID,
			&u.MainAddressID,
		)
	if err != nil {
		return u, apperror.Wrap(err)
	}
	u.DateOfBirth = toTimePtr(dateOfBirth)

	return u, nil
}

func (r *userRepositoryPostgres) IsPhoneNumberUsed(ctx context.Context, phoneNumber string) (bool, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"phoneNumber": phoneNumber,
	}
	query := `
		SELECT id 
		FROM users 
		WHERE phone_number = @phoneNumber 
			AND deleted_at IS NULL
	`

	var id int64
	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, apperror.Wrap(err)
	}

	return true, nil
}

func (r *userRepositoryPostgres) SetMainAddressByID(ctx context.Context, addressID int64, userID int64) error {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		UPDATE users 
			SET main_address_id = @addressID
		WHERE id = @userID
	`
	args := pgx.NamedArgs{
		"addressID": addressID,
		"userID":    userID,
	}

	_, err := queryRunner.ExecContext(ctx, query, args)
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}
