package router

import (
	"gin_app/app/handler/api"
	"gin_app/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/", api.ApiHandler)
		apiGroup.GET("/generate-jwt", api.GenerateJwtHandler)
		apiGroup.POST("/validator", api.ValidatorHandler)
	}
	apiJwtGroup := engine.Group("/api").Use(middleware.JwtMiddleware)
	{
		apiJwtGroup.GET("/auth-jwt", api.AuthJwtHandler)
	}
}
