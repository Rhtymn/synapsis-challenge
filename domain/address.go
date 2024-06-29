package domain

import "context"

type UserAddress struct {
	ID          int64
	Name        string
	PhoneNumber string
	Address     string
	Coordinate  Coordinate
	User        User
}

type UserAddressRepository interface {
	GetByID(ctx context.Context, id int64) (UserAddress, error)
	Add(ctx context.Context, ua UserAddress) (UserAddress, error)
}
