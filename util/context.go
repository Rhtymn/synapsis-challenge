package util

import (
	"context"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
)

func GetAccountIDFromContext(ctx context.Context) (int64, error) {
	val := ctx.Value(constants.ContextAccountID)
	id, ok := val.(int64)
	if !ok {
		return 0, apperror.NewTypeAssertionFailed(id, val)
	}
	return id, nil
}

func GetPermissionFromContext(ctx context.Context) (int64, error) {
	val := ctx.Value(constants.ContextPermission)
	permission, ok := val.(int64)
	if !ok {
		return 0, apperror.NewTypeAssertionFailed(permission, val)
	}
	return permission, nil
}
