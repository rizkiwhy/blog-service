package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

func RedisConnection() (redisClient *redis.Client, err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), DB: 0,
	})

	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Error().Err(err).Msg("[database][RedisConnection] Failed to connect to Redis")
		return
	}

	log.Info().Str("pong", pong).Msg("[database][RedisConnection] Connected to Redis")

	return
}