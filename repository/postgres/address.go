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

type userAddressRepositoryPostgres struct {
	db *sql.DB
}

func NewUserAddressRepository(db *sql.DB) *userAddressRepositoryPostgres {
	return &userAddressRepositoryPostgres{
		db: db,
	}
}

func (r *userAddressRepositoryPostgres) Add(ctx context.Context, ua domain.UserAddress) (domain.UserAddress, error) {
	userAddress := domain.UserAddress{}
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		INSERT INTO user_addresses(id_user, name, phone_number, address, latitude, longitude)
			VALUES(@userId, @name, @phoneNumber, @address, @lat, @lon)
		RETURNING ` + constants.UserAddressColumns + `
	`
	args := pgx.NamedArgs{
		"userId":      ua.User.ID,
		"name":        ua.Name,
		"phoneNumber": ua.PhoneNumber,
		"address":     ua.Address,
		"lat":         ua.Coordinate.Latitude,
		"lon":         ua.Coordinate.Longitude,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&userAddress.ID,
			&userAddress.Name,
			&userAddress.PhoneNumber,
			&userAddress.Address,
			&userAddress.Coordinate.Latitude,
			&userAddress.Coordinate.Longitude,
			&userAddress.User.ID,
		)
	if err != nil {
		return userAddress, apperror.Wrap(err)
	}

	return userAddress, nil
}

func (r *userAddressRepositoryPostgres) GetByID(ctx context.Context, id int64) (domain.UserAddress, error) {
	var userAddress domain.UserAddress
	queryRunner := util.GetQueryRunner(ctx, r.db)
	query := `
		SELECT ` + constants.UserAddressColumns + `
		FROM user_addresses 
		WHERE id = @id
			AND deleted_at IS NULL
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&userAddress.ID,
			&userAddress.Name,
			&userAddress.PhoneNumber,
			&userAddress.Address,
			&userAddress.Coordinate.Latitude,
			&userAddress.Coordinate.Longitude,
			&userAddress.User.ID,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return userAddress, apperror.NewNotFound(err, "user address not found")
		}
		return userAddress, apperror.Wrap(err)
	}
	return userAddress, nil
}
