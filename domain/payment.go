package domain

import "context"

type Payment struct {
	ID          int64
	FileURL     string
	Transaction Transaction
}

type PaymentRepository interface {
	Add(ctx context.Context, payment Payment) (Payment, error)
}
