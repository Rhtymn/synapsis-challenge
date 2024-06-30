package service

import (
	"context"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
)

type productService struct {
	productRepository domain.ProductRepository
}

type ProductServiceOpts struct {
	Product domain.ProductRepository
}

func NewProductService(opts ProductServiceOpts) *productService {
	return &productService{
		productRepository: opts.Product,
	}
}

func (s *productService) GetAll(ctx context.Context, query domain.ProductQuery) ([]domain.Product, domain.PageInfo, error) {
	products, err := s.productRepository.GetAll(ctx, query)
	if err != nil {
		return nil, domain.PageInfo{}, apperror.Wrap(err)
	}

	pageInfo, err := s.productRepository.GetPageInfo(ctx, query)
	if err != nil {
		return nil, domain.PageInfo{}, apperror.Wrap(err)
	}

	pageInfo.ItemsPerPage = int(query.Limit)
	if query.Limit == 0 {
		pageInfo.ItemsPerPage = len(products)
	}

	if pageInfo.ItemsPerPage == 0 {
		pageInfo.PageCount = 0
	} else {
		pageInfo.PageCount = (int(pageInfo.ItemCount) + pageInfo.ItemsPerPage - 1) / pageInfo.ItemsPerPage
	}

	return products, pageInfo, nil
}
