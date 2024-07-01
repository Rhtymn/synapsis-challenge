package dto

import (
	"fmt"

	"github.com/Rhtymn/synapsis-challenge/apperror"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ResponseError(appErr *apperror.AppError) Response {
	return Response{
		Message: appErr.Message,
		Data:    nil,
	}
}

func ResponseCreated(domain string, data any) Response {
	return Response{
		Message: fmt.Sprintf("%s created", domain),
		Data:    data,
	}
}

func ResponseOK(data any) Response {
	return Response{
		Message: "ok",
		Data:    data,
	}
}
