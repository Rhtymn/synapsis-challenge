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

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Add(ctx context.Context, t domain.Transaction) (domain.Transaction, error) {
	ts := domain.Transaction{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		INSERT INTO transactions(
			invoice, status, num_of_items, subtotal, 
			shipment_fee, total_fee, address, 
			latitude, longitude, phone_number,
			id_shop, id_user, id_shipment_method, id_payment_method)
		VALUES(@inv, @status, @numOfItems, @subtotal,
			@shipmentFee, @totalFee, @address,
			@latitude, @longitude, @phoneNumber,
			@idShop, @idUser, @idShipmentMethod, @idPaymentMethod	
		) RETURNING ` + constants.TransactionColumns + `
	`
	args := pgx.NamedArgs{
		"inv": t.Invoice, "status": t.Status, "numOfItems": t.NumOfItems, "subtotal": t.SubTotal,
		"shipmentFee": t.ShipmentFee, "totalFee": t.TotalFee, "address": t.Address.Address,
		"latitude": t.Address.Coordinate.Latitude, "longitude": t.Address.Coordinate.Longitude, "phoneNumber": t.Address.PhoneNumber,
		"idShop": t.Shop.ID, "idUser": t.User.ID, "idShipmentMethod": t.ShipmentMethod.ID, "idPaymentMethod": t.PaymentMethod.ID,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&ts.ID, &ts.Invoice, &ts.Status, &ts.NumOfItems, &ts.SubTotal,
			&ts.ShipmentFee, &ts.TotalFee, &ts.Address.Address,
			&ts.Address.Coordinate.Latitude, &ts.Address.Coordinate.Longitude, &ts.Address.PhoneNumber,
			&ts.Shop.ID, &ts.User.ID, &ts.ShipmentMethod.ID, &ts.PaymentMethod.ID)
	if err != nil {
		return ts, apperror.Wrap(err)
	}
	return ts, nil
}

func (r *transactionRepository) UpdateStatus(ctx context.Context, transactionID int64, status string) error {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		UPDATE transactions SET status = @status WHERE id = @id
	`
	args := pgx.NamedArgs{
		"status": status,
		"id":     transactionID,
	}

	_, err := queryRunner.ExecContext(ctx, query, args)
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}

func (r *transactionRepository) GetByInvoice(ctx context.Context, invoice string) (domain.Transaction, error) {
	ts := domain.Transaction{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		SELECT ` + constants.TransactionColumns + ` 
		FROM transactions
		WHERE invoice = @invoice 
			AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"invoice": invoice,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&ts.ID, &ts.Invoice, &ts.Status, &ts.NumOfItems, &ts.SubTotal,
			&ts.ShipmentFee, &ts.TotalFee, &ts.Address.Address,
			&ts.Address.Coordinate.Latitude, &ts.Address.Coordinate.Longitude, &ts.Address.PhoneNumber,
			&ts.Shop.ID, &ts.User.ID, &ts.ShipmentMethod.ID, &ts.PaymentMethod.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ts, apperror.NewNotFound(err, "transaction not found")
		}
		return ts, apperror.Wrap(err)
	}
	return ts, nil
}
