package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRedis struct {
	client *redis.Client
}

func NewCacheRedis(client *redis.Client) *CacheRedis {
	return &CacheRedis{client: client}
}

func (c *CacheRedis) Store(ctx context.Context, key string, value []byte, exp time.Duration) error {
	return c.client.Set(ctx, key, value, exp).Err()
}

func (c *CacheRedis) Get(ctx context.Context, key string) (val []byte, found bool, err error) {
	val, err = c.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return val, true, nil
}

func (c *CacheRedis) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

func (c *CacheRedis) Close() error {
	return c.client.Close()
}
