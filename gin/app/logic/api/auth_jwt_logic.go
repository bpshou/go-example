package api

import (
	"github.com/gin-gonic/gin"
)

func AuthJwtLogic(c *gin.Context) gin.H {
	return gin.H{
		"message": "Auth Jwt Success",
		"code":    200,
		"data": gin.H{
			"subject":  c.GetString("subject"),
			"audience": c.GetStringSlice("audience"),
		},
	}
}
