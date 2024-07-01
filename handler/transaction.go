package handler

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/dto"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionSrc domain.TransactionService
	domain         string
}

type TransactionHandlerOpts struct {
	Transaction domain.TransactionService
	Domain      string
}

func NewTransactionHandler(opts TransactionHandlerOpts) *TransactionHandler {
	return &TransactionHandler{
		transactionSrc: opts.Transaction,
		domain:         opts.Domain,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var req dto.CreateTransactionRequestDTO
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.NewBadRequest(err, "request body validation failed"))
		ctx.Abort()
		return
	}

	transaction, err := h.transactionSrc.CreateTransaction(ctx, req.ToCreateTransactionRequest())
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseOK(dto.NewTransactionResponse(transaction)))
}

func (h *TransactionHandler) PayTransaction(ctx *gin.Context) {
	var form dto.PaymentRequestDTO
	contentLength := ctx.Request.Header.Get("Content-Length")
	i, err := strconv.Atoi(contentLength)
	if err != nil {
		ctx.Error(apperror.NewInternal(err))
		ctx.Abort()
		return
	}

	if i > constants.MaxImageSize {
		ctx.Error(apperror.NewImageSizeExceeded("500kb"))
		ctx.Abort()
		return
	}

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.Error(apperror.NewBadRequest(err, "form validation failed"))
		ctx.Abort()
		return
	}

	var file *multipart.File
	fileType := form.File.Header.Get("Content-Type")
	if fileType != constants.ImageJpeg && fileType != constants.ImageJpg && fileType != constants.ImagePng {
		ctx.Error(apperror.NewRestrictredFileType(constants.ImageJpeg, constants.ImageJpg, constants.ImagePng))
		ctx.Abort()
		return
	}
	f, err := form.File.Open()
	if err != nil {
		ctx.Error(apperror.NewInternal(err))
		ctx.Abort()
		return
	}
	file = &f
	defer f.Close()

	err = h.transactionSrc.PayTransaction(ctx, form.Invoice, *file)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseOK(nil))
}
