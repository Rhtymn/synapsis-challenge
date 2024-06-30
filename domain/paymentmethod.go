package domain

import "context"

type PaymentMethod struct {
	ID   int64
	Name string
}

type PaymentMethodRepository interface {
	GetByID(ctx context.Context, id int64) (PaymentMethod, error)
}
