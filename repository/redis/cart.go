package redis

import (
	"context"
	"fmt"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/redis/go-redis/v9"
)

type cartRepositoryRedis struct {
	rc *redis.Client
}

func NewCartRepositoryRedis(rc *redis.Client) *cartRepositoryRedis {
	return &cartRepositoryRedis{
		rc: rc,
	}
}

func (r *cartRepositoryRedis) Add(ctx context.Context, accountID int64, ci domain.CartItem) error {
	cmd := r.rc.HSet(ctx,
		fmt.Sprintf("cart:%d", accountID),
		ci.Product.ID,
		fmt.Sprintf("%d:%s:%d:%d:%d:%s", ci.Product.ID, ci.Product.Name, ci.Amount, ci.TotalPrice, ci.Shop.ID, ci.Shop.ShopName),
	)
	_, err := cmd.Result()
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}

func (r *cartRepositoryRedis) GetAll(ctx context.Context, accountID int64) ([]domain.CartItem, error) {
	cart := []domain.CartItem{}
	mapCmd := r.rc.HGetAll(ctx, fmt.Sprintf("cart:%d", accountID))
	res, err := mapCmd.Result()
	if err != nil {
		return nil, apperror.Wrap(err)
	}

	for _, v := range res {
		i, err := StringToCartItem(v)
		if err != nil {
			return nil, apperror.Wrap(err)
		}
		cart = append(cart, i)
	}

	return cart, nil
}

func (r *cartRepositoryRedis) GetByID(ctx context.Context, accountID int64, productID int64) (domain.CartItem, error) {
	cart := domain.CartItem{}
	stringCmd := r.rc.HGet(ctx, fmt.Sprintf("cart:%d", accountID), fmt.Sprintf("%d", productID))
	res, err := stringCmd.Result()
	if err != nil {
		if err == redis.Nil {
			return cart, apperror.NewNotFound(err, "cart item not found")
		}
		return cart, apperror.Wrap(err)
	}

	ci, err := StringToCartItem(res)
	if err != nil {
		return cart, apperror.Wrap(err)
	}
	return ci, nil
}

func (r *cartRepositoryRedis) Delete(ctx context.Context, accountID, productID int64) error {
	intCmd := r.rc.HDel(ctx, fmt.Sprintf("cart:%d", accountID), fmt.Sprintf("%d", productID))
	_, err := intCmd.Result()
	if err != nil {
		return apperror.Wrap(err)
	}
	return nil
}
