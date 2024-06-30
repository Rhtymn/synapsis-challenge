package redis

import (
	"strconv"
	"strings"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
)

func StringToCartItem(s string) (domain.CartItem, error) {
	arr := strings.Split(s, ":")
	productID, err := ToInt64(arr[0])
	if err != nil {
		return domain.CartItem{}, err
	}

	amount, err := ToInt64(arr[2])
	if err != nil {
		return domain.CartItem{}, err
	}

	totalPrice, err := ToInt64(arr[3])
	if err != nil {
		return domain.CartItem{}, err
	}

	shopID, err := ToInt64(arr[4])
	if err != nil {
		return domain.CartItem{}, err
	}

	return domain.CartItem{
		Product: domain.Product{
			ID:   int64(productID),
			Name: arr[1],
		},
		Shop: domain.Shop{
			ID:       int64(shopID),
			ShopName: arr[5],
		},
		Amount:     amount,
		TotalPrice: totalPrice,
	}, nil
}

func ToInt64(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, apperror.Wrap(err)
	}
	return i, nil
}
