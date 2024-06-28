package dto

import "github.com/Rhtymn/synapsis-challenge/apperror"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ResponseError(appErr *apperror.AppError) Response {
	return Response{
		Message: appErr.Message,
		Data: nil,
	}
}