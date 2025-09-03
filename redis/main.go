package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	err := rds.Set(ctx, "test", "test-data", 10*time.Second).Err()
	if err != nil {
		slog.Error("redis set ", "error", err)
	}
	log.Printf("redis set success")

	result, err := rds.Get(ctx, "test").Result()
	if err != nil {
		slog.Error("redis get ", "error", err)
	}
	log.Printf("redis get success: %s", result)

	ttl, err := rds.TTL(ctx, "test").Result()
	if err != nil {
		slog.Error("redis ttl ", "error", err)
	}
	log.Printf("redis ttl success: %s", ttl)
}
