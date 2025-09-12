package api

import (
	"gin_app/app/contants"
	"gin_app/app/core/auth"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GenerateJwtLogic() gin.H {
	j := auth.NewJwt(auth.Jwt[contants.CustomClaims]{
		Issuer:     viper.GetString("Jwt.Issuer"),
		SigningKey: []byte(viper.GetString("Jwt.SigningKey")),
	})
	token, err := j.GenerateJwtTokenWithClaims(&contants.CustomClaims{
		Uid:              10001,
		Source:           "login",
		Subject:          "gin-login",
		Audience:         []string{"andy"},
		RegisteredClaims: j.GetRegisteredClaimsDefault(time.Hour * 24),
	})
	if err != nil {
		return gin.H{
			"code":    500,
			"message": "Generate Jwt Failed",
		}
	}
	return gin.H{
		"code":    0,
		"message": "Generate Jwt Success",
		"data":    token,
	}
}
