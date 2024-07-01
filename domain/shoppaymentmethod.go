package domain

import "context"

type ShopPaymentMethodRepository interface {
	IsSupportPaymentMethod(ctx context.Context, shopID int64, paymentMethodID int64) (bool, error)
}
