package middleware

import (
	"gin_app/app/core/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogMiddleware(c *gin.Context) {
	log.InitLogger(&lumberjack.Logger{
		Filename:   viper.GetString("Log.Filename"), // 日志文件的位置
		MaxSize:    viper.GetInt("Log.MaxSize"),     // 文件最大尺寸（以MB为单位）
		MaxAge:     viper.GetInt("Log.MaxAge"),      // 保留旧文件的最大天数
		MaxBackups: viper.GetInt("Log.MaxBackups"),  // 保留的最大旧文件数量
		LocalTime:  viper.GetBool("Log.LocalTime"),  // 使用本地时间创建时间戳
		Compress:   viper.GetBool("Log.Compress"),   // 是否压缩/归档旧文件
	})
	c.Next()
}
