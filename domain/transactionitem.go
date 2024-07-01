package domain

import "context"

type TransactionItem struct {
	ID          int64
	Amount      int64
	TotalPrice  int64
	Transaction Transaction
	Product     Product
}

type TransactionItemRepository interface {
	Add(ctx context.Context, ti TransactionItem) (TransactionItem, error)
}
