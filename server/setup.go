package server

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/handler"
	"github.com/Rhtymn/synapsis-challenge/middleware"
	"github.com/gin-gonic/gin"
)

type ServerOpts struct {
	AccountHandler     *handler.AccountHandler
	UserHandler        *handler.UserHandler
	ProductHandler     *handler.ProductHandler
	CartHandler        *handler.CartHandler
	TransactionHandler *handler.TransactionHandler

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

	productGroup := apiV1Group.Group("/products")
	productGroup.GET(".", opts.ProductHandler.GetAll)

	cartGroup := apiV1Group.Group("/carts",
		opts.Authenticator,
		middleware.Authorization(constants.USER_PERMISSION),
	)
	cartGroup.POST(".", opts.CartHandler.AddToCart)
	cartGroup.GET(".", opts.CartHandler.GetCart)
	cartGroup.DELETE("/:id", opts.CartHandler.DeleteCartItem)

	transactionGroup := apiV1Group.Group("/transactions",
		opts.Authenticator,
	)
	transactionGroup.POST(".",
		middleware.Authorization(constants.USER_PERMISSION),
		opts.TransactionHandler.CreateTransaction,
	)
	transactionGroup.POST("/payments",
		middleware.Authorization(constants.USER_PERMISSION),
		opts.TransactionHandler.PayTransaction,
	)

	return router
}
