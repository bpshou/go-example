package api

import (
	"gin_app/app/logic/api"
	"gin_app/app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidatorHandler(c *gin.Context) {
	var req types.ValidatorReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	resp, err := api.ValidatorLogic(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    resp,
	})
}
