package database

import (
	"context"
	"erp/internal/config"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return client, client.Ping(ctx).Err()
}
