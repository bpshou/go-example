package main

import (
	"log/slog"
	"os"

	"github.com/google/uuid"
)

func main() {
	InitSlog()
	slog.Info("hello world", "name", "John", "age", 30)
	slog.Error("error !!!", "error", "error message")
}

func InitSlog() {
	f, err := os.OpenFile("./slog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	// 在handler上加上uuid属性
	handler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}).WithAttrs([]slog.Attr{
		slog.String("uuid", uuid.NewString()),
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
