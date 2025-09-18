package {{.ServicePackage}}

import (
	"{{.LogicImportPackage}}"

	"github.com/gin-gonic/gin"
)

func {{title .ServiceName}}Handler(c *gin.Context) {
	c.JSON(200, {{.ServicePackage}}.{{title .ServiceName}}Logic(c))
}
