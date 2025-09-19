package api

import (
	"gin_app/app/core/rds"
	"gin_app/app/types"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidatorLogic(c *gin.Context, req *types.ValidatorReq) (types.ValidatorResp, error) {
	redisRes := rds.Rds.Set(c.Request.Context(), "key:name", req.Name, time.Second*10)
	if redisRes.Err() != nil {
		return types.ValidatorResp{}, redisRes.Err()
	}
	return types.ValidatorResp{}, nil
}
