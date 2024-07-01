package handler

import (
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
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

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	var form dto.UpdateProfileRequest
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
	if form.Photo != nil {
		fileType := form.Photo.Header.Get("Content-Type")
		if fileType != constants.ImageJpeg && fileType != constants.ImageJpg && fileType != constants.ImagePng {
			ctx.Error(apperror.NewRestrictredFileType(constants.ImageJpeg, constants.ImageJpg, constants.ImagePng))
			ctx.Abort()
			return
		}
		f, err := form.Photo.Open()
		if err != nil {
			ctx.Error(apperror.NewInternal(err))
			ctx.Abort()
			return
		}
		file = &f
		defer f.Close()
	}

	dob, err := time.Parse("2006-01-02", form.DateOfBirth)
	if err != nil {
		ctx.Error(apperror.NewInternal(err))
		ctx.Abort()
		return
	}

	user, err := h.userSrv.UpdateProfile(ctx, domain.UserProfile{
		Name:        form.Name,
		Photo:       file,
		DateOfBirth: dob,
		Gender:      form.Gender,
		PhoneNumber: form.PhoneNumber,
	})
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, dto.NewUserResponse(user))
}
