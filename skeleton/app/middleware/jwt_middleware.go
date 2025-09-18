package middleware

import (
	"gin_app/app/contants"
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

	// 请求头获取token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	tokenSlice := strings.Split(token, " ")
	if len(tokenSlice) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	token = tokenSlice[1]

	// 创建jwt实例
	jwt := auth.NewJwt(auth.Jwt[contants.CustomClaims]{
		Issuer:     viper.GetString("Jwt.Issuer"),
		SigningKey: []byte(viper.GetString("Jwt.SigningKey")),
	})

	var claims contants.CustomClaims
	_, err := jwt.ParseJwtTokenWithClaims(token, &claims)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("uid", claims.Uid)
	c.Next()
}
