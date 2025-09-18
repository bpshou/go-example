package api

import (
	"gin_app/app/logic/api"

	"github.com/gin-gonic/gin"
)

func ApiHandler(c *gin.Context) {
	c.JSON(200, api.ApiLogic())
}
