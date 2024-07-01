package middleware

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/dto"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors[0].Err

		apperr, ok := err.(*apperror.AppError)
		if ok {
			ctx.AbortWithStatusJSON(
				getHttpStatus(apperr),
				dto.ResponseError(apperr),
			)
			return
		}

		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			dto.ResponseError(apperror.Wrap(err).(*apperror.AppError)),
		)
	}
}

func getHttpStatus(err *apperror.AppError) int {
	switch err.Code {
	case apperror.CodeBadRequest:
		return http.StatusBadRequest
	case apperror.CodeNotFound:
		return http.StatusNotFound
	case apperror.CodeAlreadyExists:
		return http.StatusBadRequest
	case apperror.CodeUnauthorized:
		return http.StatusUnauthorized
	case apperror.CodeForbidden:
		return http.StatusForbidden
	case apperror.CodeAlreadyVerified:
		return http.StatusBadRequest
	case apperror.CodeInvalidToken:
		return http.StatusUnauthorized
	case apperror.CodeUnimplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusInternalServerError
	}
}
