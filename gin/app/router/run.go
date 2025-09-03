package router

import (
	"gin_app/app/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	// 全局中间件注册
	engine.Use(middleware.LogMiddleware)
	// 注册路由
	RegisterRouters(engine)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	engine.Run("0.0.0.0:2020")
}
