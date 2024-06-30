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

type paymentRepositoryPostgres struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *paymentRepositoryPostgres {
	return &paymentRepositoryPostgres{
		db: db,
	}
}

func (r *paymentRepositoryPostgres) Add(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	p := domain.Payment{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		INSERT INTO payments(file_url, id_transaction)
			VALUES(@fileURL, @transactionID)
		RETURNING ` + constants.PaymentColumns + `
	`
	args := pgx.NamedArgs{
		"fileURL":       payment.FileURL,
		"transactionID": payment.Transaction.ID,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&p.ID, &p.FileURL, &p.Transaction.ID)
	if err != nil {
		return p, apperror.Wrap(err)
	}
	return p, nil
}
