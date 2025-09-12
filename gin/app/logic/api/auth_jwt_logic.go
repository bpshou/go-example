package api

import (
	"github.com/gin-gonic/gin"
)

func AuthJwtLogic(c *gin.Context) gin.H {
	return gin.H{
		"code":    0,
		"message": "Auth Jwt Success",
		"data": gin.H{
			"uid": c.GetInt64("uid"),
		},
	}
}
