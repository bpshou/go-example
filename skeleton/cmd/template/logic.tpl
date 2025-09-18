package {{.ServicePackage}}

import (
	"github.com/gin-gonic/gin"
)

func {{title .ServiceName}}Logic(c *gin.Context) gin.H {
	return gin.H{
		"code":    0,
		"message": "Success",
		"data":    nil,
	}
}
