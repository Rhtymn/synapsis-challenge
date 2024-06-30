package service

import (
	"context"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
)

type cartService struct {
	cartRepository    domain.CartRepositoryRedis
	productRepository domain.ProductRepository
}

type CartServiceOpts struct {
	Cart    domain.CartRepositoryRedis
	Product domain.ProductRepository
}

func NewCartService(opts CartServiceOpts) *cartService {
	return &cartService{
		cartRepository:    opts.Cart,
		productRepository: opts.Product,
	}
}

func (s *cartService) Add(ctx context.Context, ci domain.CartItem) error {
	accountID, err := util.GetAccountIDFromContext(ctx)
	if err != nil {
		return apperror.Wrap(err)
	}

	product, err := s.productRepository.GetByID(ctx, ci.Product.ID)
	if err != nil {
		return apperror.Wrap(err)
	}

	ci.Product.Name = product.Name
	ci.TotalPrice = int(product.Price) * ci.Amount
	ci.Shop.ID = product.Shop.ID
	ci.Shop.ShopName = product.Shop.ShopName

	c, err := s.cartRepository.GetByID(ctx, accountID, product.ID)
	if err != nil {
		if !apperror.IsErrorCode(err, apperror.CodeNotFound) {
			return apperror.Wrap(err)
		}
	}

	if c.Product.ID == ci.Product.ID {
		ci.Amount += c.Amount
		ci.TotalPrice = int(product.Price) * ci.Amount
	}

	err = s.cartRepository.Add(ctx, accountID, ci)
	if err != nil {
		return apperror.Wrap(err)
	}

	return nil
}

func (s *cartService) GetAll(ctx context.Context) ([]domain.CartItem, error) {
	accountID, err := util.GetAccountIDFromContext(ctx)
	if err != nil {
		return nil, apperror.Wrap(err)
	}

	cart, err := s.cartRepository.GetAll(ctx, accountID)
	if err != nil {
		return nil, apperror.Wrap(err)
	}

	return cart, nil
}

func (s *cartService) DeleteCartItem(ctx context.Context, productID int64) error {
	accountID, err := util.GetAccountIDFromContext(ctx)
	if err != nil {
		return apperror.Wrap(err)
	}

	_, err = s.cartRepository.GetByID(ctx, accountID, productID)
	if err != nil {
		return apperror.Wrap(err)
	}

	err = s.cartRepository.Delete(ctx, accountID, productID)
	if err != nil {
		return apperror.Wrap(err)
	}

	return nil
}
