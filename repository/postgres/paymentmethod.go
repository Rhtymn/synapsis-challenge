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

type paymentMethodRepositoryPostgres struct {
	db *sql.DB
}

func NewPaymentMethodRepository(db *sql.DB) *paymentMethodRepositoryPostgres {
	return &paymentMethodRepositoryPostgres{
		db: db,
	}
}

func (r *paymentMethodRepositoryPostgres) GetByID(ctx context.Context, id int64) (domain.PaymentMethod, error) {
	p := domain.PaymentMethod{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		SELECT` + constants.PaymentMethodColumns + `
		FROM payment_methods 
		WHERE id = @id 
			AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&p.ID, &p.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, apperror.NewNotFound(err, "payment method not found")
		}
		return p, apperror.Wrap(err)
	}

	return p, nil
}
