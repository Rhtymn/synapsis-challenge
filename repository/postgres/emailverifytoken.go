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

type emailVerifyTokenRepositoryPostgres struct {
	db *sql.DB
}

func NewEmailVerifyTokenRepository(db *sql.DB) *emailVerifyTokenRepositoryPostgres {
	return &emailVerifyTokenRepositoryPostgres{
		db: db,
	}
}

func (r *emailVerifyTokenRepositoryPostgres) GetById(ctx context.Context, id int64) (domain.EmailVerifyToken, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	emailVerifyToken := domain.EmailVerifyToken{}
	query := `
		SELECT ` + constants.EmailVerifyTokenColumns + `
		FROM email_verify_tokens 
		WHERE id = @id 
			AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&emailVerifyToken.ID,
			&emailVerifyToken.Token,
			&emailVerifyToken.ExpiredAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return emailVerifyToken, apperror.NewNotFound(err, "email verify token not found")
		}
		return emailVerifyToken, apperror.Wrap(err)
	}
	return emailVerifyToken, nil
}

func (r *emailVerifyTokenRepositoryPostgres) GetByTokenStr(ctx context.Context, tokenStr string) (domain.EmailVerifyToken, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	emailVerifyToken := domain.EmailVerifyToken{}
	query := `
		SELECT ` + constants.EmailVerifyTokenColumns + `
		FROM email_verify_tokens 
		WHERE token = @token
			AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"token": tokenStr,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&emailVerifyToken.ID,
			&emailVerifyToken.Token,
			&emailVerifyToken.ExpiredAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return emailVerifyToken, apperror.NewNotFound(err, "email verify token not found")
		}
		return emailVerifyToken, apperror.Wrap(err)
	}
	return emailVerifyToken, nil
}

func (r *emailVerifyTokenRepositoryPostgres) GetByTokenStrAndLock(ctx context.Context, tokenStr string) (domain.EmailVerifyToken, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	emailVerifyToken := domain.EmailVerifyToken{}
	query := `
		SELECT ` + constants.EmailVerifyTokenColumns + `
		FROM email_verify_tokens 
		WHERE token = @token
			AND deleted_at IS NULL
		FOR UPDATE
	`
	args := pgx.NamedArgs{
		"token": tokenStr,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&emailVerifyToken.ID,
			&emailVerifyToken.Token,
			&emailVerifyToken.ExpiredAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return emailVerifyToken, apperror.NewNotFound(err, "email verify token not found")
		}
		return emailVerifyToken, apperror.Wrap(err)
	}
	return emailVerifyToken, nil
}

func (r *emailVerifyTokenRepositoryPostgres) Add(ctx context.Context, e domain.EmailVerifyToken) (domain.EmailVerifyToken, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	emailVerifyToken := domain.EmailVerifyToken{}
	query := `
		INSERT INTO email_verify_tokens(token, expired_at, id_account)
			VALUES(@token, @expiredAt, @accountId)
		RETURNING ` + constants.EmailVerifyTokenColumns + `
	`
	args := pgx.NamedArgs{
		"token":     e.Token,
		"expiredAt": e.ExpiredAt,
		"accountId": e.Account.ID,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&emailVerifyToken.ID,
			&emailVerifyToken.Token,
			&emailVerifyToken.ExpiredAt,
		)
	if err != nil {
		return emailVerifyToken, apperror.Wrap(err)
	}
	return emailVerifyToken, nil
}

func (r *emailVerifyTokenRepositoryPostgres) SoftDeleteByToken(ctx context.Context, token string) error {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		UPDATE email_verify_tokens SET deleted_at = NOW()
			WHERE token = @token
				AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"token": token,
	}

	_, err := queryRunner.ExecContext(ctx, query, args)
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}

func (r *emailVerifyTokenRepositoryPostgres) SoftDeleteByAccountID(ctx context.Context, id int64) error {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		UPDATE email_verify_tokens SET deleted_at = NOW()
			WHERE id_account = @accountID 
				AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"accountID": id,
	}

	_, err := queryRunner.ExecContext(ctx, query, args)
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}
