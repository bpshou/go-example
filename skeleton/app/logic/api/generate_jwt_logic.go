package api

import (
	"gin_app/app/contants"
	"gin_app/app/core/auth"
	"gin_app/app/types"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GenerateJwtLogic(c *gin.Context, req *types.GenerateJwtReq) (types.GenerateJwtResp, error) {
	j := auth.NewJwt(auth.Jwt[contants.CustomClaims]{
		Issuer:     viper.GetString("Jwt.Issuer"),
		SigningKey: []byte(viper.GetString("Jwt.SigningKey")),
	})
	token, err := j.GenerateJwtTokenWithClaims(&contants.CustomClaims{
		Uid:              req.Uid,
		Source:           "login",
		Subject:          "gin-login",
		Audience:         []string{"andy"},
		RegisteredClaims: j.GetRegisteredClaimsDefault(time.Hour * 24),
	})
	if err != nil {
		return types.GenerateJwtResp{}, err
	}
	return types.GenerateJwtResp{
		Token: token,
	}, nil
}
