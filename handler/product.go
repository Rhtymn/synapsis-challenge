package handler

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/dto"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productSrv domain.ProductService
	domain     string
}

type ProductHandlerOpts struct {
	Product domain.ProductService
	Domain  string
}

func NewProductHandler(opts ProductHandlerOpts) *ProductHandler {
	return &ProductHandler{
		productSrv: opts.Product,
		domain:     opts.Domain,
	}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	var query dto.GetProductQuery
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "query validation failed"))
		ctx.Abort()
		return
	}

	products, pageInfo, err := h.productSrv.GetAll(ctx, query.ToProductQuery())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.NewProductsResponse(products, pageInfo))
}
