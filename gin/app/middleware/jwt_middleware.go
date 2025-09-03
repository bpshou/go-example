package middleware

import (
	"gin_app/app/core/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func JwtMiddleware(c *gin.Context) {
	if !viper.IsSet("Jwt.Issuer") {
		viper.Set("Jwt.Issuer", "gin-app")
	}
	if !viper.IsSet("Jwt.SigningKey") {
		viper.Set("Jwt.SigningKey", "gin-app")
	}

	// 创建jwt实例
	jwt := auth.NewJwt(auth.Jwt{
		Issuer:     viper.GetString("Jwt.Issuer"),
		SigningKey: []byte(viper.GetString("Jwt.SigningKey")),
	})
	// 请求头获取token
	token := c.GetHeader("Authorization")
	token = strings.Split(token, " ")[1]

	subject, audience, err := jwt.ParseJwtToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("subject", subject)
	c.Set("audience", audience)
	c.Next()
}
