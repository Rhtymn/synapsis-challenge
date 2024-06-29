package handler

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/dto"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountSrv domain.AccountService
	domain     string
}

type AccountHandlerOpts struct {
	Account domain.AccountService
	Domain  string
}

func NewAccountHandler(opts AccountHandlerOpts) *AccountHandler {
	return &AccountHandler{
		accountSrv: opts.Account,
		domain:     opts.Domain,
	}
}

func (h *AccountHandler) Register(ctx *gin.Context) {
	var req dto.AccountRegisterRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "request body validation failed"))
		ctx.Abort()
		return
	}

	var params dto.AccountTypeParams
	err = ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "params validation failed"))
		ctx.Abort()
		return
	}

	_, err = h.accountSrv.Register(ctx, req.ToCredentials(params.Type))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseCreated(h.domain, nil))
}

func (h *AccountHandler) Login(ctx *gin.Context) {
	var req dto.AccountLoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "request body validation failed"))
		ctx.Abort()
		return
	}

	token, err := h.accountSrv.Login(ctx, req.ToCredentials())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.NewAuthTokenResponse(token))
}
