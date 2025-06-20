package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
