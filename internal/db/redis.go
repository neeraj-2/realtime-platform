package db

import (
	"context"
	"realtime-platform/internal/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func ConnectRedis(cfg *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	_, err := client.Ping(Ctx).Result()
	if err != nil {
		config.Log.Fatal("Failed to connect to Redis: " + err.Error())
	}

	config.Log.Info("Successfully connected to Redis")
	return client
}