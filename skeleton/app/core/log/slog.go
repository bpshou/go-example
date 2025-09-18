package log

import (
	"log/slog"

	"github.com/google/uuid"
	"gopkg.in/natefinch/lumberjack.v2"
)

var SlogLevel = new(slog.LevelVar)

// lumberjack.Logger{
// 	Filename:   "./lumberjack.log", // 日志文件的位置
// 	MaxSize:    1,                  // 文件最大尺寸（以MB为单位）
// 	MaxAge:     1,                  // 保留旧文件的最大天数
// 	MaxBackups: 3,                  // 保留的最大旧文件数量
// 	LocalTime:  true,               // 使用本地时间创建时间戳
// 	Compress:   false,              // 是否压缩/归档旧文件
// }

func InitLogger(log *lumberjack.Logger) {
	SlogLevel.Set(slog.LevelInfo)

	// 在handler上加上uuid属性
	handler := slog.NewJSONHandler(log, &slog.HandlerOptions{
		AddSource: true,
		Level:     SlogLevel,
	}).WithAttrs([]slog.Attr{
		slog.String("uuid", uuid.NewString()),
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
