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

type accountRepositoryPostgres struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *accountRepositoryPostgres {
	return &accountRepositoryPostgres{
		db: db,
	}
}

func (r *accountRepositoryPostgres) GetById(ctx context.Context, id int64) (domain.AccountWithCredentials, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	a := domain.AccountWithCredentials{}
	args := pgx.NamedArgs{
		"id": id,
	}
	query := `
		SELECT ` + constants.AccountWithCredentialColumns + `
		FROM accounts 
		WHERE id = @id 
			AND deleted_at IS NULL
	`

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&a.Account.ID,
			&a.Account.Email,
			&a.Account.EmailVerified,
			&a.Password,
			&a.Account.AccountType,
			&a.Account.ProfileSet,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return a, apperror.NewNotFound(err, "account not found")
		}
		return a, apperror.Wrap(err)
	}

	return a, nil
}

func (r *accountRepositoryPostgres) GetByIdAndLock(ctx context.Context, id int64) (domain.AccountWithCredentials, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	a := domain.AccountWithCredentials{}
	args := pgx.NamedArgs{
		"id": id,
	}
	query := `
		SELECT ` + constants.AccountWithCredentialColumns + `
		FROM accounts 
		WHERE id = @id 
			AND deleted_at IS NULL
		FOR UPDATE
	`

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&a.Account.ID,
			&a.Account.Email,
			&a.Account.EmailVerified,
			&a.Password,
			&a.Account.AccountType,
			&a.Account.ProfileSet,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return a, apperror.NewNotFound(err, "account not found")
		}
		return a, apperror.Wrap(err)
	}

	return a, nil
}

func (r *accountRepositoryPostgres) GetByEmail(ctx context.Context, email string) (domain.AccountWithCredentials, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	a := domain.AccountWithCredentials{}
	args := pgx.NamedArgs{
		"email": email,
	}
	query := `
		SELECT ` + constants.AccountWithCredentialColumns + `
		FROM accounts 
		WHERE email = @email 
			AND deleted_at IS NULL
	`

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&a.Account.ID,
			&a.Account.Email,
			&a.Account.EmailVerified,
			&a.Password,
			&a.Account.AccountType,
			&a.Account.ProfileSet,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return a, apperror.NewNotFound(err, "account not found")
		}
		return a, apperror.Wrap(err)
	}

	return a, nil
}

func (r *accountRepositoryPostgres) GetByEmailAndLock(ctx context.Context, email string) (domain.AccountWithCredentials, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	a := domain.AccountWithCredentials{}
	args := pgx.NamedArgs{
		"email": email,
	}
	query := `
		SELECT ` + constants.AccountWithCredentialColumns + `
		FROM accounts 
		WHERE email = @email 
			AND deleted_at IS NULL
		FOR UPDATE
	`

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&a.Account.ID,
			&a.Account.Email,
			&a.Account.EmailVerified,
			&a.Password,
			&a.Account.AccountType,
			&a.Account.ProfileSet,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return a, apperror.NewNotFound(err, "account not found")
		}
		return a, apperror.Wrap(err)
	}

	return a, nil
}

func (r *accountRepositoryPostgres) Add(ctx context.Context, a domain.AccountWithCredentials) (domain.Account, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	account := domain.Account{}
	query := `
		INSERT INTO accounts(email, password, account_type)
			VALUES(@email, @password, @accountType)
		RETURNING ` + constants.AccountColumns + `
	`
	args := pgx.NamedArgs{
		"email":       a.Account.Email,
		"password":    a.Password,
		"accountType": a.Account.AccountType,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&account.ID,
			&account.Email,
			&account.EmailVerified,
			&account.AccountType,
			&account.ProfileSet,
		)
	if err != nil {
		return account, apperror.Wrap(err)
	}
	return account, nil
}

func (r *accountRepositoryPostgres) VerifyEmailById(ctx context.Context, id int64) error {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"id": id,
	}
	query := `
		UPDATE accounts 
			SET email_verified = true, updated_at = NOW()
		WHERE id = @id 
			AND deleted_at IS NULL
	`
	_, err := queryRunner.ExecContext(ctx, query, args)
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}

func (r *accountRepositoryPostgres) ProfileSetById(ctx context.Context, id int64) error {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	args := pgx.NamedArgs{
		"id": id,
	}
	query := `
		UPDATE accounts
			SET profile_set = true, updated_at = NOW()
		WHERE id = @id
			AND deleted_at IS NULL
	`

	_, err := queryRunner.ExecContext(ctx, query, args)
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}

func (r *accountRepositoryPostgres) IsEmailUsed(ctx context.Context, email string) (bool, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		SELECT id
		FROM accounts
		WHERE email = @email
			AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"email": email,
	}

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
