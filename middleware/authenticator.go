package middleware

import (
	"strings"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/gin-gonic/gin"
)

func Authenticator(jwtProvider util.JWTProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authVal := ctx.GetHeader("Authorization")

		if authVal == "" {
			ctx.Error(apperror.NewUnauthorized(nil, "unauthorized"))
			ctx.Abort()
			return
		}

		tokens := strings.Fields(authVal)
		if len(tokens) != 2 {
			ctx.Error(apperror.NewInvalidToken(nil))
			ctx.Abort()
			return
		}
		if tokens[0] != "Bearer" {
			ctx.Error(apperror.NewInvalidToken(nil))
			ctx.Abort()
			return
		}

		token := tokens[1]

		claims, err := jwtProvider.VerifyToken(token)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Set(constants.ContextAccountID, claims.AccountID)
		ctx.Next()
	}
}
