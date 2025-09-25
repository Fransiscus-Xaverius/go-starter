package repository

import (
	"context"
	"time"
)

type CacheFile struct {
}

func (c CacheFile) Ping(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c CacheFile) Store(ctx context.Context, key string, value []byte, exp time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (c CacheFile) Get(ctx context.Context, key string) ([]byte, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CacheFile) Close() error {
	//TODO implement me
	panic("implement me")
}
