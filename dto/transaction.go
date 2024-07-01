package dto

import (
	"mime/multipart"

	"github.com/Rhtymn/synapsis-challenge/domain"
)

type CreateTransactionRequestDTO struct {
	ProductID        int64 `json:"product_id" binding:"required,numeric"`
	ShipmentMethodID int64 `json:"shipment_method_id" binding:"required,numeric"`
	PaymentMethodID  int64 `json:"payment_method_id" binding:"required,numeric"`
	AddressID        int64 `json:"address_id" binding:"required,numeric"`
}

type PaymentRequestDTO struct {
	Invoice string                `form:"invoice" binding:"required"`
	File    *multipart.FileHeader `form:"file" binding:"required"`
}

type TransactionResponse struct {
	ID          int64  `json:"id"`
	Invoice     string `json:"invoice"`
	Status      string `json:"status"`
	NumOfItems  int64  `json:"num_of_items"`
	Subtotal    int64  `json:"subtotal"`
	ShipmentFee int64  `json:"shipment_fee"`
	TotalFee    int64  `json:"total_fee"`
}

func (c *CreateTransactionRequestDTO) ToCreateTransactionRequest() domain.CreateTransactionRequest {
	return domain.CreateTransactionRequest{
		ProductID:        c.ProductID,
		ShipmentMethodID: c.ShipmentMethodID,
		PaymentMethodID:  c.PaymentMethodID,
		AddressID:        c.AddressID,
	}
}

func NewTransactionResponse(t domain.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:          t.ID,
		Invoice:     t.Invoice,
		Status:      t.Status,
		NumOfItems:  t.NumOfItems,
		Subtotal:    t.SubTotal,
		ShipmentFee: t.ShipmentFee,
		TotalFee:    t.TotalFee,
	}
}
