package repository

import (
	"context"
	"database/sql"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/jackc/pgx/v5"
)

type shopPaymentMethodRepositoryPostgres struct {
	db *sql.DB
}

func NewShopPaymentMethodRepository(db *sql.DB) *shopPaymentMethodRepositoryPostgres {
	return &shopPaymentMethodRepositoryPostgres{
		db: db,
	}
}

func (r *shopPaymentMethodRepositoryPostgres) IsSupportPaymentMethod(ctx context.Context, shopID int64, paymentMethodID int64) (bool, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		SELECT id FROM shop_payment_methods
			WHERE id_shop = @shopID 
				AND id_payment_method = @paymentMethodID
				AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"shopID":          shopID,
		"paymentMethodID": paymentMethodID,
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
