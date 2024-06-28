package server

import (
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	return router
}
