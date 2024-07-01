package dto

import "github.com/Rhtymn/synapsis-challenge/domain"

type AddToCartRequest struct {
	ProductID int64 `json:"product_id" binding:"required,numeric"`
	Amount    int   `json:"amount" binding:"required,numeric,min=1"`
}

type CartResponse struct {
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Amount      int64   `json:"amount"`
	TotalPrice  int64   `json:"total_price"`
	Shop        ShopDTO `json:"shop,omitempty"`
}

func NewCartResponse(ci domain.CartItem) CartResponse {
	return CartResponse{
		ProductID:   ci.Product.ID,
		ProductName: ci.Product.Name,
		Amount:      int64(ci.Amount),
		TotalPrice:  int64(ci.TotalPrice),
		Shop:        NewShopDTO(ci.Shop),
	}
}

func GetCartResponse(ci []domain.CartItem) Response {
	cr := []CartResponse{}
	for i := 0; i < len(ci); i++ {
		cr = append(cr, NewCartResponse(ci[i]))
	}
	return Response{
		Message: "successfully fetch cart",
		Data:    cr,
	}
}
