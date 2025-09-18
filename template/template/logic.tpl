package {{.logicPackage}}

import (
	"github.com/gin-gonic/gin"
)

func {{title .serviceName}}Logic(c *gin.Context) gin.H {
	return gin.H{
		"code":    0,
		"message": "Success",
		"data":    nil,
	}
}
