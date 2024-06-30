package handler

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/dto"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartSrv domain.CartServiceRedis
	domain  string
}

type CartHandlerOpts struct {
	Cart   domain.CartServiceRedis
	Domain string
}

func NewCartHandler(opts CartHandlerOpts) *CartHandler {
	return &CartHandler{
		cartSrv: opts.Cart,
		domain:  opts.Domain,
	}
}

func (h *CartHandler) AddToCart(ctx *gin.Context) {
	var req dto.AddToCartRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "request body validation failed"))
		ctx.Abort()
		return
	}

	err = h.cartSrv.Add(ctx, domain.CartItem{
		Product: domain.Product{
			ID: req.ProductID,
		},
		Amount: req.Amount,
	})
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseOK(nil))
}

func (h *CartHandler) GetCart(ctx *gin.Context) {
	cart, err := h.cartSrv.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.GetCartResponse(cart))
}

func (h *CartHandler) DeleteCartItem(ctx *gin.Context) {
	var params dto.IDParams
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "params validation failed"))
		ctx.Abort()
		return
	}

	err = h.cartSrv.DeleteCartItem(ctx, params.ID)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Status(http.StatusNoContent)
}
