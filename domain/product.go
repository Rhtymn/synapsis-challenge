package domain

import "context"

const (
	SortProductByName  = "name"
	SortProductByPrice = "price"
)

type Product struct {
	ID          int64
	Name        string
	Slug        string
	PhotoURL    *string
	Price       int64
	Description *string
	Stock       int64
	Category    Category
	Shop        Shop
}

type ProductQuery struct {
	Page         int64
	Limit        int64
	SortBy       string
	SortType     string
	CategorySlug string
	Search       string
}

type ProductRepository interface {
	GetByID(ctx context.Context, id int64) (Product, error)
	GetAll(ctx context.Context, query ProductQuery) ([]Product, error)
	GetPageInfo(ctx context.Context, query ProductQuery) (PageInfo, error)
}

type ProductService interface {
	GetAll(ctx context.Context, query ProductQuery) ([]Product, PageInfo, error)
}
