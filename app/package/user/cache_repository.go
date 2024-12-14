package user

import (
	"encoding/json"
	"rizkiwhy-blog-service/package/user/model"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

type CacheRepository interface {
	SetJWTPayload(request model.SetJWTPayloadRequest) (err error)
}

type CacheRepositoryImpl struct {
	RedisClient *redis.Client
}

func NewCacheRepository(redisClient *redis.Client) CacheRepository {
	return &CacheRepositoryImpl{
		RedisClient: redisClient,
	}
}

func (c *CacheRepositoryImpl) SetJWTPayload(request model.SetJWTPayloadRequest) (err error) {
	request.KeyJWTPayload()
	request.ValueJWTPayload()

	valueJSON, err := json.Marshal(request.Value)
	if err != nil {
		log.Error().Err(err).Msg("[SetJWTPayload] Failed to marshal value to JSON")
		return
	}

	err = c.RedisClient.Set(c.RedisClient.Context(), request.Key, valueJSON, request.Exp).Err()
	if err != nil {
		log.Error().Err(err).Msg("[SetJWTPayload] Failed to set value in Redis")
	}

	return
}
