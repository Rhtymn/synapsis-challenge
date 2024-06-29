package middleware

import (
	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/gin-gonic/gin"
)

func Authorization(permissions ...int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		permission, err := util.GetPermissionFromContext(ctx)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		for i := 0; i < len(permissions); i++ {
			if permission == permissions[i] {
				ctx.Next()
				return
			}
		}

		ctx.Error(apperror.NewForbidden(nil, "forbidden"))
		ctx.Abort()
		return
	}
}
