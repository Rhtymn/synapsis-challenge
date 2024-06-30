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

type transactionRepositoryPostgres struct {
	db *sql.DB
}

func NewTransactionItemRepository(db *sql.DB) *transactionRepositoryPostgres {
	return &transactionRepositoryPostgres{
		db: db,
	}
}

func (r *transactionRepositoryPostgres) Add(ctx context.Context, ti domain.TransactionItem) (domain.TransactionItem, error) {
	transactionItem := domain.TransactionItem{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		INSERT INTO transaction_items(amount, total_price, id_transaction, id_product)
			VALUES(@amount, @totalPrice, @transactionID, @productID)
		RETURNING ` + constants.TransactionItemColumns + `
	`
	args := pgx.NamedArgs{
		"amount":        ti.Amount,
		"totalPrice":    ti.TotalPrice,
		"transactionID": ti.Transaction.ID,
		"productID":     ti.Product.ID,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&transactionItem.ID,
			&transactionItem.Amount,
			&transactionItem.TotalPrice,
			&transactionItem.Transaction.ID,
			&transactionItem.Product.ID,
		)
	if err != nil {
		return transactionItem, apperror.Wrap(err)
	}
	return transactionItem, nil
}
