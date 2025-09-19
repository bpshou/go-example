package api

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func ApiLogic() gin.H {
	slog.Info("api logic", "data", "success")

	return gin.H{
		"code":    0,
		"message": "Hello, World!",
		"data":    "data",
	}
}
