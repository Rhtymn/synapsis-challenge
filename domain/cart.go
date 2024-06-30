package domain

import "context"

type CartItem struct {
	Product    Product
	Shop       Shop
	Amount     int
	TotalPrice int
}

type CartRepositoryRedis interface {
	Add(ctx context.Context, accountID int64, ci CartItem) error
	GetAll(ctx context.Context, accountID int64) ([]CartItem, error)
	GetByID(ctx context.Context, accountID int64, productID int64) (CartItem, error)
	Delete(ctx context.Context, accountID, productID int64) error
}

type CartServiceRedis interface {
	GetAll(ctx context.Context) ([]CartItem, error)
	Add(ctx context.Context, ci CartItem) error
	DeleteCartItem(ctx context.Context, productID int64) error
}
