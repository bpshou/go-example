package api

import (
	"gin_app/app/logic/api"

	"github.com/gin-gonic/gin"
)

func AuthJwtHandler(c *gin.Context) {
	c.JSON(200, api.AuthJwtLogic(c))
}
