package handler

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSrv domain.UserService
	domain  string
}

type UserHandlerOpts struct {
	User   domain.UserService
	Domain string
}

func NewUserHandler(opts UserHandlerOpts) *UserHandler {
	return &UserHandler{
		userSrv: opts.User,
		domain:  opts.Domain,
	}
}

func (h *UserHandler) AddAddress(ctx *gin.Context) {
	var req dto.CreateAddressRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "request body validation failed"))
		ctx.Abort()
		return
	}

	userAddress, err := h.userSrv.AddAddress(ctx, req.ToUserAddressDomain())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseCreated("user address", dto.NewUserAddressResponse(userAddress)))
}

func (h *UserHandler) UpdateMainAddress(ctx *gin.Context) {
	var params dto.UpdateMainAddressParams
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "params validation failed"))
		ctx.Abort()
		return
	}

	err = h.userSrv.UpdateMainAddress(ctx, params.AddressID)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Status(http.StatusNoContent)
}
