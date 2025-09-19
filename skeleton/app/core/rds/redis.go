package rds

import (
	"github.com/redis/go-redis/v9"
)

var Rds *redis.Client

func NewRedis(addr string, password string, db int) *redis.Client {
	Rds = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return Rds
}
