package domain

import "context"

type ShopShipmentMethodRepository interface {
	IsSupportShipmentMethod(ctx context.Context, shopID int64, shipmentMethodID int64) (bool, error)
}
