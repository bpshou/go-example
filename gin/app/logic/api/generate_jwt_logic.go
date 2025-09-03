package api

import (
	"gin_app/app/core/auth"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GenerateJwtLogic() gin.H {
	jwt := auth.NewJwt(auth.Jwt{
		Issuer:     viper.GetString("Jwt.Issuer"),
		SigningKey: []byte(viper.GetString("Jwt.SigningKey")),
	})
	token, err := jwt.GenerateJwtToken("login", []string{"zhangsan", "12"}, time.Hour*24)
	if err != nil {
		return gin.H{
			"message": "Generate Jwt Failed",
			"code":    500,
		}
	}
	return gin.H{
		"message": "Generate Jwt Success",
		"code":    200,
		"data":    token,
	}
}
