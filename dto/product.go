package dto

import (
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/domain"
)

type GetProductQuery struct {
	Page         int64  `form:"page" binding:"numeric,omitempty,min=1"`
	Limit        int64  `form:"limit" binding:"numeric,omitempty,min=1"`
	SortBy       string `form:"sort_by" binding:"omitempty,oneof=name price"`
	SortType     string `form:"sort_type" binding:"omitempty,oneof=asc desc"`
	CategorySlug string `form:"category_slug"`
	Search       string `form:"search"`
}

type ProductsResponseWithPageInfo struct {
	Products []ProductDTO `json:"products"`
	PageInfo PageInfo     `json:"page_info"`
}

type ProductDTO struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Slug        string  `json:"slug,omitempty"`
	PhotoURL    *string `json:"photo_url,omitempty"`
	Price       int64   `json:"price,omitempty"`
	Description *string `json:"description,omitempty"`
	Stock       int64   `json:"stock,omitempty"`
	Shop        ShopDTO `json:"shop,omitempty"`
}

func (q *GetProductQuery) ToProductQuery() domain.ProductQuery {
	page := q.Page
	sortBy := q.SortBy
	sortType := q.SortType
	if q.Page == 0 || q.Limit == 0 {
		page = 1
	}
	if q.SortBy == "" {
		sortBy = domain.SortProductByName
	}
	if q.SortType == "" {
		sortType = constants.SortASC
	}
	return domain.ProductQuery{
		Page:         page,
		Limit:        q.Limit,
		SortBy:       sortBy,
		SortType:     sortType,
		CategorySlug: q.CategorySlug,
		Search:       q.Search,
	}
}

func NewProductsResponse(products []domain.Product, p domain.PageInfo) Response {
	productsDTO := []ProductDTO{}
	for i := 0; i < len(products); i++ {
		productsDTO = append(productsDTO, ProductDTO{
			ID:          products[i].ID,
			Name:        products[i].Name,
			Slug:        products[i].Slug,
			PhotoURL:    products[i].PhotoURL,
			Price:       products[i].Price,
			Description: products[i].Description,
			Stock:       products[i].Stock,
			Shop:        NewShopDTO(products[i].Shop),
		})
	}

	return Response{
		Message: "successfully fetch products",
		Data: ProductsResponseWithPageInfo{
			Products: productsDTO,
			PageInfo: NewPageInfoResponse(p),
		},
	}
}
