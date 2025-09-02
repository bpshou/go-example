package main

import (
	"log/slog"

	"github.com/google/uuid"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	InitLogger()
	slog.Info("hello world", "name", "John", "age", 30) // 这行日志会输出

	LogLevel.Set(slog.LevelError) // 这样就可以在运行时更新日志等级

	slog.Info("hello world", "name", "Rose", "age", 18) // 这行日志不会输出

	slog.Error("error !!!", "error", "error message") // 这行日志会输出

	// 写入100000行日志，测试文件切割功能
	for i := 0; i < 100000; i++ {
		slog.Error("error test message !!!", "data", "data connection error")
	}
}

var LogLevel = new(slog.LevelVar)

func InitLogger() {
	log := lumberjack.Logger{
		Filename:   "./lumberjack.log", // 日志文件的位置
		MaxSize:    1,                  // 文件最大尺寸（以MB为单位）
		MaxAge:     1,                  // 保留旧文件的最大天数
		MaxBackups: 3,                  // 保留的最大旧文件数量
		LocalTime:  true,               // 使用本地时间创建时间戳
		Compress:   false,              // 是否压缩/归档旧文件
	}

	LogLevel.Set(slog.LevelInfo)

	// 在handler上加上uuid属性
	handler := slog.NewJSONHandler(&log, &slog.HandlerOptions{
		AddSource: true,
		Level:     LogLevel,
	}).WithAttrs([]slog.Attr{
		slog.String("uuid", uuid.NewString()),
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
