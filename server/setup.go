package server

import (
	"net/http"

	"github.com/Rhtymn/synapsis-challenge/handler"
	"github.com/gin-gonic/gin"
)

type ServerOpts struct {
	AccountHandler *handler.AccountHandler

	CorsHandler  gin.HandlerFunc
	ErrorHandler gin.HandlerFunc
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

	return router
}
