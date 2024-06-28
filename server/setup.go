package server

import (
	"github.com/gin-gonic/gin"
)

type ServerOpts struct {
	CorsHandler gin.HandlerFunc
}

func SetupServer(opts ServerOpts) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(
		gin.Recovery(),
		opts.CorsHandler,
	)

	return router
}
