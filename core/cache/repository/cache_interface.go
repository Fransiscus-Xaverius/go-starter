package repository

import (
	"context"
	"time"
)

//go:generate mockgen -destination=mocks/mock_CacheInterface.go -package=mocks . CacheInterface
type CacheInterface interface {
	Ping(ctx context.Context) error
	Store(ctx context.Context, key string, value []byte, exp time.Duration) error
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Close() error
}
