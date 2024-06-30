package domain

import (
	"context"
	"mime/multipart"
)

const (
	WaitingForPayment      = "waiting for payment"
	WaitingForConfirmation = "waiting for confirmation"
)

type Transaction struct {
	ID             int64
	Invoice        string
	Status         string
	NumOfItems     int64
	SubTotal       int64
	ShipmentFee    int64
	TotalFee       int64
	Address        UserAddress
	Shop           Shop
	User           User
	ShipmentMethod ShipmentMethod
	PaymentMethod  PaymentMethod
}

type CreateTransactionRequest struct {
	ProductID        int64
	ShipmentMethodID int64
	PaymentMethodID  int64
	AddressID        int64
}

type TransactionRepository interface {
	Add(ctx context.Context, t Transaction) (Transaction, error)
	UpdateStatus(ctx context.Context, transactionID int64, status string) error
	GetByInvoice(ctx context.Context, invoice string) (Transaction, error)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, t CreateTransactionRequest) (Transaction, error)
	PayTransaction(ctx context.Context, invoice string, file multipart.File) error
}
