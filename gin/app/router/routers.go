package router

import (
	"gin_app/app/handler/api"
	"gin_app/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine) {
	router := engine.Group("/api")
	{
		router.GET("/", api.ApiHandler)
		router.GET("/generate-jwt", api.GenerateJwtHandler)
		router.Use(middleware.JwtMiddleware).GET("/auth-jwt", api.AuthJwtHandler)
	}
}
