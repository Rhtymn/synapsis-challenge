package repository

import (
	"context"
	"database/sql"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/jackc/pgx/v5"
)

type shopShipmentMethodRepositoryPostgres struct {
	db *sql.DB
}

func NewShopShipmentMethodRepository(db *sql.DB) *shopShipmentMethodRepositoryPostgres {
	return &shopShipmentMethodRepositoryPostgres{
		db: db,
	}
}

func (r *shopShipmentMethodRepositoryPostgres) IsSupportShipmentMethod(ctx context.Context, shopID int64, shipmentMethodID int64) (bool, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		SELECT id FROM shop_shipment_methods
			WHERE id_shop = @shopID 
				AND id_shipment_method = @shipmentMethodID
				AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"shopID":           shopID,
		"shipmentMethodID": shipmentMethodID,
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
