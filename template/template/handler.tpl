package {{.handlerPackage}}

import (
	"{{.logicImportPackage}}"

	"github.com/gin-gonic/gin"
)

func {{title .serviceName}}Handler(c *gin.Context) {
	c.JSON(200, {{.logicPackage}}.{{title .serviceName}}Logic(c))
}
