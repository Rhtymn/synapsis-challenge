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

type shipmentMethodRepositoryPostgres struct {
	db *sql.DB
}

func NewShipmentMethodRepository(db *sql.DB) *shipmentMethodRepositoryPostgres {
	return &shipmentMethodRepositoryPostgres{
		db: db,
	}
}

func (r *shipmentMethodRepositoryPostgres) GetByID(ctx context.Context, id int64) (domain.ShipmentMethod, error) {
	p := domain.ShipmentMethod{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		SELECT` + constants.ShipmentMethodColumns + `
		FROM shipment_methods 
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
			return p, apperror.NewNotFound(err, "shipment method not found")
		}
		return p, apperror.Wrap(err)
	}

	return p, nil
}
