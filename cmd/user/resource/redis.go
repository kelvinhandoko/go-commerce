package resource

import (
	"context"
	"fmt"
	"go-commerce/config"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
	})

	// test connection
	ctx := context.Background()
	pingResult, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Redis connected:", pingResult)

	return RedisClient
}
