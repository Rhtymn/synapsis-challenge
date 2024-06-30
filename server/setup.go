package server

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/handler"
	"github.com/Rhtymn/synapsis-challenge/middleware"
	"github.com/gin-gonic/gin"
)

type ServerOpts struct {
	AccountHandler *handler.AccountHandler
	UserHandler    *handler.UserHandler

	CorsHandler  gin.HandlerFunc
	ErrorHandler gin.HandlerFunc

	Authenticator gin.HandlerFunc
}

func SetupServer(opts ServerOpts) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(
		gin.Recovery(),
		opts.CorsHandler,
		opts.ErrorHandler,
	)

	apiV1Group := router.Group("/api/v1")
	apiV1Group.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	authGroup := apiV1Group.Group("/auth")
	authGroup.POST("/register/:type", opts.AccountHandler.Register)
	authGroup.POST("/login", opts.AccountHandler.Login)
	authGroup.POST("/verify-token", opts.AccountHandler.VerifyEmail)
	authGroup.GET(
		"/verify-token",
		opts.Authenticator,
		middleware.Authorization(constants.SELLER_PERMISSION, constants.USER_PERMISSION),
		opts.AccountHandler.GetVerifyEmailToken,
	)
	authGroup.GET("/check-verify-token", opts.AccountHandler.CheckVerifyEmailToken)

	userGroup := apiV1Group.Group("/users",
		opts.Authenticator,
		middleware.Authorization(constants.USER_PERMISSION),
	)
	userGroup.POST("/addresses", opts.UserHandler.AddAddress)
	userGroup.PATCH("/addresses/:address_id/main", opts.UserHandler.UpdateMainAddress)
	userGroup.PUT(".", opts.UserHandler.UpdateProfile)

	return router
}
