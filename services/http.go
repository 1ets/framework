package services

import (
	"framework/app/adapters/servers"

	"github.com/gin-gonic/gin"
)

// Global Middleware setup
func HttpMiddleware(middleware *gin.Engine) {
	middleware.Use(gin.Recovery())
}

func HttpRouter(route *gin.Engine) {
	route.POST("example", servers.HttpPostExample)
}
