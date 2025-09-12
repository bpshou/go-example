package router

import (
	"gin_app/app/middleware"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.New() // gin.Default()
	// debug
	gin.SetMode(gin.DebugMode)
	// 控制台和日志文件双输出
	gin.DefaultWriter = io.MultiWriter(os.Stdout, middleware.GetLumberjackLogger())
	// 全局中间件注册
	engine.Use(gin.Recovery(), gin.Logger(), middleware.LogMiddleware)
	// 注册路由
	RegisterRouters(engine)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	engine.Run("0.0.0.0:2020")
}
